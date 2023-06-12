package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/helper"
	"github.com/vinoMamba/pharos-go/internal/router"
)

func TestCreateSession(t *testing.T) {
	r := router.New()

	email := "vinoTest@qq.com"
	code := helper.GernerateDigits(6)
	q := database.NewQuery()
	c := context.Background()
	if _, err := q.CreateValidaitonCode(c, queries.CreateValidaitonCodeParams{
		Email: email,
		Code:  code,
	}); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	body := gin.H{
		"email": email,
		"code":  code,
	}

	bytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		strings.NewReader(string(bytes)),
	)

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var responseBody struct {
		jwt string
	}

	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
