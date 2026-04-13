package middleware

import (
	"errors"
	"net/http"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func AuthorizationMiddleware(db tools.DatabaseInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var AccountID string = r.URL.Query().Get("account_id")
			var token = r.Header.Get("Authorization")

			if AccountID == "" || token == "" {
				log.Error(UnAuthorizedError)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			var loginDetails *tools.LoginDetails
			loginDetails = db.GetUserLoginDetails(AccountID)

			if loginDetails == nil || (token != (*loginDetails).AuthToken) {
				log.Error(UnAuthorizedError)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
