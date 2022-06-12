package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.ShowError(c, 422, err.Error())
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.ShowSuccess(c, gin.H{
			"token": token,
			"user":  user,
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {

	// 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.ShowError(c, 422, err.Error())
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.ShowSuccess(c, gin.H{
			"token": token,
			"user":  user,
		})
	}
}
