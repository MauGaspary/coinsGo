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
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.AccountID == "" || req.Password == "" {
		log.Error("Dados de registro inválidos: ", err)
		api.RequestErrorHandler(w, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Erro ao gerar hash da senha: ", err)
		api.InternalErrorHandler(w)
		return
	}

	_, err = h.DB.CreateAccount(r.Context(), req.AccountID)
	if err != nil {
		log.Error("Erro ao criar a conta: ", err)
		api.InternalErrorHandler(w)
		return
	}

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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Conta criada com sucesso!"})
	log.Info("Conta criada")
}