package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/database"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	AccountID string `json:"account_id"`
	Password  string `json:"password"`
}

func (h *AccountHandlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// 1. Lemos o JSON que o usuário enviou no corpo (body) da requisição
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.AccountID == "" || req.Password == "" {
		api.RequestErrorHandler(w, err)
		return
	}

	// 2. Geramos o hash da senha usando o bcrypt (nunca salvamos em texto plano)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Erro ao gerar hash da senha: ", err)
		api.InternalErrorHandler(w)
		return
	}

	// 3. Criamos a conta no banco (com saldo 0.0 por padrão)
	_, err = h.DB.CreateAccount(r.Context(), req.AccountID)
	if err != nil {
		log.Error("Erro ao criar a conta: ", err)
		api.InternalErrorHandler(w)
		return
	}

	// 4. Criamos os detalhes de login com o hash
	params := database.CreateLoginDetailParams{
		AccountID:    req.AccountID,
		PasswordHash: string(hashedPassword),
	}
	_, err = h.DB.CreateLoginDetail(r.Context(), params)
	if err != nil {
		log.Error("Erro ao salvar detalhes de login: ", err)
		api.InternalErrorHandler(w)
		return
	}

	// 5. Retornamos sucesso (201 Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Conta criada com sucesso!"})
}