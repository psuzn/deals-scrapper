package api

import (
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

var Logger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
	Logger: log.New(),
})
