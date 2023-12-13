package api

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func lb(str string) string {
	return fmt.Sprintf("%s\n", str)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	BuildRouter().ServeHTTP(rr, req)
	return rr
}

func TestLivenessApi(t *testing.T) {
	t.Run("should return up", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/health/liveness", nil)
		response := executeRequest(request)

		require.Equal(t, 200, response.Code)
		require.Equal(t, lb(`{"status":"UP"}`), response.Body.String())
	})
}

func TestReadinessApi(t *testing.T) {
	t.Run("should return up", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/health/readiness", nil)

		response := executeRequest(request)

		require.Equal(t, 200, response.Code)
		require.Equal(t, lb(`{"status":"UP"}`), response.Body.String())
	})
}
