package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	FAILED  = 1
)

//普通成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"smg":  "成功",
		"data": v,
	})
}

//普通失败返回
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"smg":  "失败",
		"data": v,
	})
}
