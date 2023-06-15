package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(c *gin.RouterGroup)
	Create(c *gin.Context)
	Destroy(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetPaged(c *gin.Context)
}
