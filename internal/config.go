package internal

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Api       Api
	serverUrl url.URL
	urls      []url.URL
}

type Api struct {
	Port int
}

type env[T int | string | url.URL] struct {
	key          string
	defaultValue T
}

var violations []string

func BuildConfig() (Config, []string) {
	violations = []string{}

	return Config{
		Api: Api{
			Port: getEnv(env[int]{key: "PORT", defaultValue: 3000}),
		},
		serverUrl: getEnv(env[url.URL]{key: "SERVER_URL"}),
		urls:      getUrls(),
	}, violations
}

func (apiConfig Api) Addr() string {
	return fmt.Sprintf(":%d", apiConfig.Port)
}

func getUrls() []url.URL {
	var rawUrls = getEnv(env[string]{key: "URLS", defaultValue: " "})
	if rawUrls == " " {
		return make([]url.URL, 0)
	}

	var urls = strings.Split(rawUrls, ";")
	var actualUrls []url.URL

	for _, entry := range urls {
		url_, err := url.ParseRequestURI(entry)
		if err == nil {
			actualUrls = append(actualUrls, *url_)
		} else {
			violations = append(violations, fmt.Sprintf("'%s' is not a valid uri", entry))
		}
	}

	return actualUrls
}

func getEnv[T int | string | url.URL](_env env[T]) T {
	var rawValue = os.Getenv(_env.key)
	var value any = *new(T)

	if rawValue == "" {
		// default value was provided
		if value.(T) != _env.defaultValue {
			return _env.defaultValue
		} else {
			violations = append(violations, fmt.Sprintf("Env '%s' was not provided", _env.key))
			return value.(T)
		}
	}

	var err error

	switch any(_env.defaultValue).(type) {
	case int:
		value, err = strconv.Atoi(rawValue)

	case string:
		value = rawValue

	case url.URL:
		var tmpUrl *url.URL
		tmpUrl, err = url.ParseRequestURI(rawValue)
		if tmpUrl != nil {
			value = *tmpUrl
		}
	}

	if err != nil {
		violations = append(violations, fmt.Sprintf(
			"Provided env '%s' value '%s' is not of type '%T'", _env.key, rawValue, value))
	}

	return value.(T)
}
