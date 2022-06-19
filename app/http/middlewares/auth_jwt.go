// Package middlewares Gin 中间件
package middlewares

import (
	"github.com/spf13/cast"
	"gohub/app/models/user"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthJWT 验证用户身份
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.ShowError(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(cast.ToString(claims.UserID))
		if userModel.ID == 0 {
			response.ShowError(c, http.StatusUnauthorized, "用户不存在")
			c.Abort()
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("user_id", userModel.GetStringID())
		c.Set("user_name", userModel.Name)
		c.Set("user", userModel)

		c.Next()
	}
}
