package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/pharos-go/internal/router"
)

func TestCreateSession(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()

	body := gin.H{
		"email": "1@qq.com",
		"code":  "123",
	}

	bytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		strings.NewReader(string(bytes)),
	)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
