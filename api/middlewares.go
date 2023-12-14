package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

var loggerMiddleware = middleware.RequestLogger(&middleware.DefaultLogFormatter{
	Logger: log.New(),
})

func setupMiddlewares(router chi.Router) {
	router.Use(loggerMiddleware)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.RequestID)
}
