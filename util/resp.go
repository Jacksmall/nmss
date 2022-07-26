package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespJson(ctx *gin.Context, code int, msg string, data []interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
