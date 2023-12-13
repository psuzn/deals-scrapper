package internal

import (
	"bytes"
	"fmt"
	"github.com/psuzn/deals-scrapper/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func submitNewAppPackage(config Config, packageName string) {
	data := []byte(fmt.Sprintf(`{ "packageName" : "%s" }`, packageName))

	response, err := http.Post(config.serverUrl.String(), api.ContentTypes.Json, bytes.NewBuffer(data))

	if err != nil {
		log.Errorf("got error '%s' for packageName '%s'", err.Error(), packageName)
	} else if response.StatusCode != 200 {
		log.Infof("response is '%d', '%s'", response.StatusCode, response.Body)
	}
}
