package main

import (
	"github.com/psuzn/deals-scrapper/api"
	"github.com/psuzn/deals-scrapper/internal"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	config, violations := internal.BuildConfig()

	if len(violations) != 0 {
		log.Fatal(violations)
	}

	go internal.ScheduleScrapping(config)

	var router = api.BuildRouter()
	log.Info("Starting api server @", config.Api.Addr())
	log.Fatal(http.ListenAndServe(config.Api.Addr(), router))
}
