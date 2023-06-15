package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/pharos-go/config"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/helper"
)

func TestCreateSession(t *testing.T) {
	config.LoadAppConfig()
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	database.Connect()
	group := r.Group("/api")
	{
		sc := SessionController{}
		sc.Register(group)
	}

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

	var u queries.User
	u, err := q.FindUserByEmail(c, email)
	if err != nil || u.ID == 0 {
		u, err = q.CreateUser(c, email)
		if err != nil {
			t.Fatal(err)
		}
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
		JWT    string `json:"jwt"`
		UserId int32  `json:"userId"`
	}

	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, u.ID, responseBody.UserId)
}
