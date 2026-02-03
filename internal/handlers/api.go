package handlers

import (
	"github.com/MauGaspary/goapi/api/internal/handlers/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.AuthotizationMiddleware)
		router.Get("/balance", GetAccountBalance)
	})
}