package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	teardown := setupTestCase(t)
	defer teardown(t)

	vc := ValidationCodeController{}
	vc.Register(r.Group("/api"))

	count1, _ := q.CountValidationCodes(c)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/validation_codes", strings.NewReader(`{"email":"1@qq.com"}`)) // 将字符串转换为 io.Reader
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, count1+1, count2)

}
