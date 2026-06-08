package handlers

import (
	"github.com/MauGaspary/goapi/internal/database"
	"github.com/MauGaspary/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux, db database.Querier) {
	r.Use(chimiddleware.StripSlashes)
	AccountHandlers := &AccountHandlers{DB: db}

	// Rota pública para criar conta (registro)
	r.Post("/register", AccountHandlers.CreateAccount)

	r.Post("/login", AccountHandlers.Login)

	// Rotas protegidas por autenticação
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.AuthorizationMiddleware(db))
		router.Get("/balance", AccountHandlers.GetAccountBalance)
	})
}
