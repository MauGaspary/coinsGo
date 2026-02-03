package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MauGaspary/coinsGo.git/api"
	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.AccountBalanceResponse{}
	var decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())