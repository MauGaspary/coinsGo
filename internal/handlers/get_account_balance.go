package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/database"
	"github.com/MauGaspary/goapi/internal/middleware"
	log "github.com/sirupsen/logrus"
)

type AccountHandlers struct {
	DB database.Querier
}

// GetAccountBalance godoc
// @Summary      Obtém o saldo de uma conta
// @Description  Endpoint para consultar o saldo da conta autenticada
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200 {object} api.AccountBalanceResponse
// @Failure      401 {object} api.Error
// @Router       /account/balance [get]
func (h *AccountHandlers) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	accountID, ok := middleware.GetAccountIDFromContext(r.Context())
	if !ok || accountID == "" {
		log.Error("account_id não encontrado no contexto da requisição")
		api.RequestErrorHandler(w, middleware.UnAuthorizedError)
		return
	}

	accountDetails, err := h.DB.GetUserAccount(r.Context(), accountID)
	if err != nil {
		log.Error("Erro ao buscar a conta no banco de dados:", err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.AccountBalanceResponse{
		Balance: accountDetails.Balance,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}