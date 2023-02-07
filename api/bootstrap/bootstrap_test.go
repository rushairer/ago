package bootstrap

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestSetupRouter(t *testing.T) {
	server := gin.Default()

	SetupServer(server)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test/alive", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
