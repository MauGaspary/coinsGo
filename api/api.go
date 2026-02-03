package api

import (
	"encoding/json"
	"net/http"
)

type AccountBalanceParams struct {
	AccountID string `schema:"account_id"`
}

type AccountBalanceResponse struct {
	Code    int     `json:"code"`
	Balance float64 `json:"balance"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeErrorResponse(w, http.StatusInternalServerError, "Unexpected Error")
	}
)
