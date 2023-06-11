package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/email"
	"github.com/vinoMamba/pharos-go/internal/helper"
)

type CreateValidationCodeParams struct {
	Email string `json:"email" binding:"required,email"`
}

func CreateValidationCode(c *gin.Context) {
	var body CreateValidationCodeParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	code := helper.GernerateDigits(6)
	q := database.NewQuery()
	vc, err := q.CreateValidaitonCode(c, queries.CreateValidaitonCodeParams{
		Email: body.Email,
		Code:  code,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "创建验证码失败"})
		return
	}

	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		c.JSON(500, gin.H{"error": "发送验证码失败"})
		return
	}
	c.JSON(200, gin.H{"message": "发送验证码成功"})
}
