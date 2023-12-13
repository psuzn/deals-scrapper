package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/psuzn/deals-scrapper/api"
	"github.com/psuzn/deals-scrapper/internal"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	var router = chi.NewRouter()
	router.Use(api.Logger)
	config, violations := internal.BuildConfig()

	if len(violations) != 0 {
		log.Error(violations)
		os.Exit(-1)
	}

	log.Info("Starting api server @", config.Api.Addr())
	err := http.ListenAndServe(config.Api.Addr(), router)

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}
