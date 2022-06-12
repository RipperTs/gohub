// Package response 统一数据返回结果
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowError 响应错误信息
func ShowError(c *gin.Context, code int, msg ...string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  defaultMessage("error", msg...),
		"data": nil,
	})
}

// ShowSuccess 响应成功信息
func ShowSuccess(c *gin.Context, data any, msg ...string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  defaultMessage("success", msg...),
		"data": data,
	})
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
