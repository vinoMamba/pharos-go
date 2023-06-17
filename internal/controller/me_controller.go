package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba/pharos-go/internal/jwt_helper"
)

type MeController struct{}

func (ctrl *MeController) Register(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.GET("/me", ctrl.Get)
}

func (ctrl *MeController) Get(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if len(auth) < 8 {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	jwtStr := auth[7:]
	if jwtStr == "" {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	t, err := jwt_helper.ParseJWT(jwtStr)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	m, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}
	user_id := m["user_id"].(int32)
	c.JSON(200, gin.H{"user_id": user_id})
}

func (ctrl *MeController) Create(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *MeController) Destory(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *MeController) Update(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *MeController) GetPaged(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}
