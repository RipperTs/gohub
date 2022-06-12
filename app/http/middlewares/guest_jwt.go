package middlewares

import (
	"gohub/pkg/jwt"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.ShowError(c, http.StatusUnauthorized, "已登录，请先退出")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
