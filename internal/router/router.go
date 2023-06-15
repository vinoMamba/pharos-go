package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config"
	"github.com/vinoMamba/pharos-go/internal/controller"
	"github.com/vinoMamba/pharos-go/internal/database"
)

func loadControllers() []controller.Controller {
	return []controller.Controller{
		&controller.SessionController{},
		&controller.ValidationCodeController{},
	}
}

func New() *gin.Engine {
	config.LoadAppConfig()
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	database.Connect()
	group := r.Group("/api")
	{
		for _, c := range loadControllers() {
			c.Register(group)
		}
	}
	return r
}
