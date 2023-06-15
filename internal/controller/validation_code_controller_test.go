package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/pharos-go/config"
	"github.com/vinoMamba/pharos-go/internal/database"
)

func TestCreateValidationCode(t *testing.T) {
	config.LoadAppConfig()
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	database.Connect()
	group := r.Group("/api")
	{
		vc := ValidationCodeController{}
		vc.Register(group)
	}

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
