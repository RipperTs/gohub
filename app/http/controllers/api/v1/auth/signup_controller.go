// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	response.ShowSuccess(c, user.IsPhoneExist(request.Phone))
}

// IsEmailExist 检查邮箱是否被注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 表单验证
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	// 返回响应结果
	response.ShowSuccess(c, user.IsEmailExist(request.Email))
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	// 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 创建数据
	userModel := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.ShowSuccess(c, gin.H{
			"token":    token,
			"userInfo": userModel,
		})
		return
	}
	response.ShowError(c, 422, "注册失败")
}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.ShowSuccess(c, gin.H{
			"token":    token,
			"userInfo": userModel,
		})
		return
	}
	response.ShowError(c, 422, "注册失败")
}
