package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/internal/email"
)

func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Println("--------")
		log.Println(err)
		log.Println("--------")
		c.JSON(500, gin.H{"error": "发送验证码失败"})
		return
	}
	c.JSON(200, gin.H{"message": "发送验证码成功"})
}
