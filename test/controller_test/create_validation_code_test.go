package controller_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/router"
)

func TestCreateValidationCode(t *testing.T) {
	r := router.New()
	c := context.Background()
	q := database.NewQuery()
	count1, _ := q.CountValidationCodes(c)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/validation_codes", strings.NewReader(`{"email":"1@qq.com"}`)) // 将字符串转换为 io.Reader
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, count1+1, count2)

}
