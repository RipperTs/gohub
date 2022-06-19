// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"
	"net/http"
)

// RegisterAPIRoutes 注册API路由
func RegisterAPIRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello " + config.GetString("app.name"),
		})
	})
	// 所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")

		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 1000 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.LimitPerRoute("20-H"), suc.IsPhoneExist)
			// 判断邮箱是否被注册
			authGroup.POST("/signup/email/exist", middlewares.LimitPerRoute("20-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			// 发送短信验证码
			authGroup.POST("/verify-codes/sms", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			// 发送邮箱验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 手机号注册用户
			authGroup.POST("/signup/using-phone", middlewares.LimitPerRoute("20-H"), suc.SignupUsingPhone)
			// 邮箱注册用户
			authGroup.POST("/signup/using-email", middlewares.LimitPerRoute("20-H"), suc.SignupUsingEmail)

			// 用户登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 密码登录
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.Use(middlewares.GuestJWT())
			{
				// 手机号重置
				authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
				// 邮箱重置
				authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)
			}
		}

	}
}
