package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println("--------")
	log.Println(body.Email)
}
