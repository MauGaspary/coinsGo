package handlers

import (
	"github.com/MauGaspary/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.AuthorizationMiddleware)
		router.Get("/balance", GetAccountBalance)
	})
}
