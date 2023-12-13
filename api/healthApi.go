package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type status string

const (
	UP   status = "UP"
	DOWN status = "DOWN"
)

type statusResponse struct {
	Status status `json:"status"`
}

func healthApi(router chi.Router) {
	router.Get("/liveness", func(w http.ResponseWriter, r *http.Request) {
		SendData(w, r, statusResponse{UP})
	})

	router.Get("/readiness", func(w http.ResponseWriter, request *http.Request) {
		SendData(w, request, statusResponse{UP})
	})
}
