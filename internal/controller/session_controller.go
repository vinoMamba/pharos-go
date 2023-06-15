package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/jwt_helper"
)

type SessionController struct{}

func (ctrl *SessionController) Register(c *gin.RouterGroup) {
	v1 := c.Group("/v1")
	v1.POST("/session", ctrl.Create)
}

func (ctrl *SessionController) Create(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	q := database.NewQuery()
	_, err := q.FindValidationCode(c, queries.FindValidationCodeParams{
		Email: body.Email,
		Code:  body.Code,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的验证码"})
		return
	}

	// FindOrCreateUser
	var u queries.User
	u, err = q.FindUserByEmail(c, body.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "请稍后再试")
		return
	}
	if u.ID == 0 {
		u, err = q.CreateUser(c, body.Email)
		if err != nil {
			c.String(http.StatusInternalServerError, "请稍后再试")
			return
		}
	}
	jwt, err := jwt_helper.GenerateJWT(1)
	if err != nil {
		c.String(http.StatusInternalServerError, "请稍后再试")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":    jwt,
		"userId": u.ID,
	})
}

func (ctrl *SessionController) Destory(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *SessionController) Update(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *SessionController) Get(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *SessionController) GetPaged(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}
