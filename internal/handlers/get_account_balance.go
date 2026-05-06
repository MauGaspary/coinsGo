package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/database"
	// "github.com/MauGaspary/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

type AccountHandlers struct {
	DB database.Querier
}

func (h *AccountHandlers) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.AccountBalanceParams{}
	var decoder *schema.Decoder= schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// var tokenDetails *database.Account
	accountDetails, err := h.DB.GetUserAccount(r.Context(), params.AccountID)
	if err != nil {
		log.Error("Erro ao buscar a conta no banco de dados:", err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.AccountBalanceResponse{
		Balance: accountDetails.Balance,
		Code: http.StatusOK,
	}


	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}