package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config"
	"github.com/vinoMamba/pharos-go/internal/controller"
	"github.com/vinoMamba/pharos-go/internal/database"
)

func New() *gin.Engine {
	config.LoadAppConfig()
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	database.Connect()
	group := r.Group("/api/v1")
	{
		group.POST("/login/dingtalk", controller.Login)
		group.POST("/validation_codes", controller.CreateValidationCode)
		group.POST("/session", controller.CreateSession)
	}
	return r
}
