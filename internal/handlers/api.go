package handlers

import (
	"github.com/MauGaspary/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/MauGaspary/goapi/internal/database"
)

func Handler(r *chi.Mux, db database.Querier) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		AccountHandlers := &AccountHandlers{DB: db}
		router.Use(middleware.AuthorizationMiddleware(db))
		router.Get("/balance", AccountHandlers.GetAccountBalance)
	})
}
