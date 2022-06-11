package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 422,
		"msg":  msg,
	})
}

func ShowSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
