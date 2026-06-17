package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/MauGaspary/goapi/api"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	AccountID string `json:"account_id"`
	Password  string `json:"password"`
}

// Login godoc
// @Summary      Faz login na conta
// @Description  Endpoint para fazer login na conta autenticada
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Dados do usuário"
// @Success      200 {object} api.LoginResponse
// @Failure      401 {object} api.Error
// @Router       /login [post]
func (h *AccountHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.AccountID == "" || req.Password == "" {
		log.Error("Dados de login inválidos")
		api.RequestErrorHandler(w, errors.New("invalid request payload"))
		return
	}

	loginDetails, err := h.DB.GetUserLoginDetails(r.Context(), req.AccountID)
	if err != nil {
		log.Error("Usuário não encontrado ou erro no banco: ", err)
		api.RequestErrorHandler(w, errors.New("invalid credentials"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginDetails.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Error("Senha incorreta")
		api.RequestErrorHandler(w, errors.New("invalid credentials"))
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Error("Variável JWT_SECRET não está configurada")
		api.InternalErrorHandler(w)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": req.AccountID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error("Erro ao assinar o token JWT: ", err)
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
