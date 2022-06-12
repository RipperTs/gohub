// Package response 统一数据返回结果
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowError 响应错误信息
func ShowError(c *gin.Context, code int, msg any) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

// ShowSuccess 响应成功信息
func ShowSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
