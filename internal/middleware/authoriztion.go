package middleware

import (
	"errors"
	"net/http"

	"github.com/MauGaspary/coinsGo.git/api"
	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func AuthotizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var AccountID string = r.URL.Query().Get("account_id")
		var token = r.Header.Get("Authorization")
		var err error

		if AccountID == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		
		var database *tools.DatabaseInterface
		database, err = tools,NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*&database).GetUserLoginDetails(AccountID)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}