package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/internal/controller"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	group := r.Group("/api/v1")
	{
		group.POST("/login/dingtalk", controller.Login)
		group.POST("/validation_codes", controller.CreateValidationCode)
	}
	return r
}
