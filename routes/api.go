// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"
)

// RegisterAPIRoutes 注册API路由
func RegisterAPIRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello GoHub!",
		})
	})
	// 所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否被注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 发送短信验证码
			authGroup.POST("/verify-codes/sms", vcc.SendUsingPhone)
			// 发送邮箱验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			// 手机号注册用户
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 邮箱注册用户
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 用户登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)

		}
	}
}
