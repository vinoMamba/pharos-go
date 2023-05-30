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
	r.GET("/api/v1/ping", controller.Ping)
	return r
}
