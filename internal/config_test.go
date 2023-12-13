package internal

import (
	"github.com/stretchr/testify/assert"
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
	})

	t.Run("correct value should be  filled", func(t *testing.T) {
		t.Setenv("PORT", "3333")
		t.Setenv("SERVER_URL", "test_url")

		config, violations := BuildConfig()

		assert.Empty(t, violations)
		assert.Equal(t, 3333, config.Api.Port)
		assert.Equal(t, "test_url", config.serverUrl)
	})

}
