// internal/api/router_test.go
package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRepositories(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/repos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
