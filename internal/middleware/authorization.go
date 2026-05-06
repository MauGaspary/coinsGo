package middleware

import (
	"errors"
	"net/http"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/database"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func AuthorizationMiddleware(db database.Querier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var AccountID string = r.URL.Query().Get("account_id")
			var token = r.Header.Get("Authorization")

			if AccountID == "" || token == "" {
				log.Error(UnAuthorizedError)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			loginDetails, err := db.GetUserLoginDetails(r.Context(),AccountID)

			if err != nil {
				log.Error("Usuário não encontrado ou erro no banco:", err)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			if token != loginDetails.AuthToken {
				log.Error(UnAuthorizedError)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
