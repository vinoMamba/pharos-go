package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
	"github.com/vinoMamba/pharos-go/internal/email"
	"github.com/vinoMamba/pharos-go/internal/helper"
)

type ValidationCodeController struct{}

func (ctrl *ValidationCodeController) Register(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	v1.POST("/validation_codes", ctrl.Create)
}

func (ctrl *ValidationCodeController) Create(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	code := helper.GernerateDigits(6)
	q := database.NewQuery()
	vc, err := q.CreateValidaitonCode(c, queries.CreateValidaitonCodeParams{
		Email: body.Email,
		Code:  code,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "创建验证码失败"})
		return
	}

	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		c.JSON(500, gin.H{"error": "发送验证码失败"})
		return
	}
	c.JSON(200, gin.H{"message": "发送验证码成功"})
}

func (ctrl *ValidationCodeController) Destory(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *ValidationCodeController) Update(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *ValidationCodeController) Get(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}

func (ctrl *ValidationCodeController) GetPaged(c *gin.Context) {
	panic("not implemented") //TODO: Implement
}
