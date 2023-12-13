package internal

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestBuildConfig(t *testing.T) {

	t.Run("should return violations when required env type is not given", func(t *testing.T) {
		t.Setenv("PORT", "12")
		_, violations := BuildConfig()

		assert.NotEmpty(t, violations)
		assert.Equal(t, "Env 'SERVER_URL' was not provided", violations[0])
	})

	t.Run("should return violations when invalid type is given", func(t *testing.T) {
		t.Setenv("PORT", "test_port")
		t.Setenv("SERVER_URL", "test_url")
		_, violations := BuildConfig()

		assert.NotEmpty(t, violations)
		assert.Equal(t, "Provided env 'PORT' value 'test_port' is not of type 'int'", violations[0])

	})

	t.Run("default value should be auto filled", func(t *testing.T) {
		t.Setenv("SERVER_URL", "test_url")
		t.Setenv("PORT", "")

		config, violations := BuildConfig()

		assert.Empty(t, violations)
		assert.Equal(t, 3000, config.Api.Port)
		assert.Equal(t, "test_url", config.serverUrl)
		assert.Equal(t, make([]url.URL, 0), config.urls)
	})

	t.Run("correct value should be  filled", func(t *testing.T) {
		t.Setenv("PORT", "3333")
		t.Setenv("SERVER_URL", "test_url")
		t.Setenv("URLS", "https://test.com;https://test2.com")

		config, violations := BuildConfig()

		assert.Empty(t, violations)
		assert.Equal(t, 3333, config.Api.Port)
		assert.Equal(t, "test_url", config.serverUrl)
		assert.Equal(t, "https://test.com", config.urls[0].String())
		assert.Equal(t, "https://test2.com", config.urls[1].String())
	})

}

func TestGetUrls(t *testing.T) {
	t.Run("should correctly parse urls", func(t *testing.T) {
		t.Setenv("URLS", "https://example.com;https://example1.com")

		urls := getUrls()

		assert.Len(t, urls, 2)
		assert.Equal(t, "https://example.com", urls[0].String())
		assert.Equal(t, "https://example1.com", urls[1].String())
	})

	t.Run("should add violation if a url is not valid url", func(t *testing.T) {
		t.Setenv("URLS", "--;https://example1.com")

		_ = getUrls()

		assert.Len(t, violations, 1)
		assert.Equal(t, "'--' is not a valid uri", violations[0])
	})

}
