package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MauGaspary/goapi/api"
	"github.com/MauGaspary/goapi/internal/database"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

type contextKey string

const AccountIDKey contextKey = "account_id"

// GetAccountIDFromContext recupera o account_id de dentro do contexto.
func GetAccountIDFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(AccountIDKey).(string)
	return val, ok
}

var UnAuthorizedError = errors.New("Invalid username or token")

func AuthorizationMiddleware(db database.Querier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token = r.Header.Get("Authorization")

			if token == "" {
				log.Error("Token de autorização ausente")
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			// Extrai apenas o token, removendo o prefixo "Bearer " (comum em APIs REST)
			tokenString := strings.TrimPrefix(token, "Bearer ")

			// Faz o parse e a validação do token JWT usando a mesma chave secreta do Login
			parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !parsedToken.Valid {
				log.Error("Token JWT inválido ou expirado: ", err)
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			// Valida e extrai o account_id de dentro do token
			claims, ok := parsedToken.Claims.(jwt.MapClaims)
			if !ok {
				log.Error("Falha ao extrair claims do token")
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			accountID, ok := claims["account_id"].(string)
			if !ok || accountID == "" {
				log.Error("account_id ausente ou inválido no token")
				api.RequestErrorHandler(w, UnAuthorizedError)
				return
			}

			// Injeta o account_id no contexto do request
			ctx := context.WithValue(r.Context(), AccountIDKey, accountID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
