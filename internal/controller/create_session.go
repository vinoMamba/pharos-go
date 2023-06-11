package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
)

type CreateSessionParams struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func CreateSession(c *gin.Context) {

	var body CreateSessionParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	q := database.NewQuery()
	_, err := q.FindValidationCode(c, queries.FindValidationCodeParams{
		Email: body.Email,
		Code:  body.Code,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的验证码"})
		return
	}

	jwt := "jwt"
	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
