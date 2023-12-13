package api

import (
	"github.com/go-chi/chi/v5"
)

func setupRoutes(router chi.Router) {
	router.Route("/health", healthApi)
}

func BuildRouter() *chi.Mux {
	var router = chi.NewRouter()
	setupMiddlewares(router)
	setupRoutes(router)

	return router
}
