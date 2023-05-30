package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"tokenInfo": gin.H{
				"tokenName":  "token",
				"tokenValue": "142a55a1-76e0-43b0-8233-4cfb5597286e",
			},
			"userInfo": gin.H{
				"id":     "1234",
				"name":   "admin",
				"avatar": "",
			},
		},
	})
}
