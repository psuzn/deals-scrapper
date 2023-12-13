package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type httpResponse[T any] struct {
	code    int
	headers map[string]string
	body    T
}

var headers = struct {
	contentType string
}{
	contentType: "Content-Type",
}

var contentTypes = struct {
	json string
}{
	json: "application/json",
}

func SendData(w http.ResponseWriter, r *http.Request, body any) {
	sendHttpRepose(w, r, httpResponse[any]{
		body:    body,
		headers: map[string]string{headers.contentType: contentTypes.json},
	},
	)
}

func sendHttpRepose(w http.ResponseWriter, r *http.Request, response httpResponse[any]) {
	if response.code == 0 {
		response.code = 200
	}

	for key, value := range response.headers {
		w.Header().Set(key, value)
	}

	err := json.NewEncoder(w).Encode(response.body)
	if err != nil {
		log.Error(err)
	}
}
