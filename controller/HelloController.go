package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (r *HelloController) Router(Engine *gin.Engine) {
	Engine.GET("/hello", r.hello)
}

func (c *HelloController) hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"mes": "成功",
	})
}
