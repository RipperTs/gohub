package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodeRequest struct {
	Captcha_id string `json:"captcha_id,omitempty" valid:"captcha_id"`
	Answer     string `json:"answer,omitempty" valid:"answer"`
}

type SmsVerifyCodeRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// 图形验证码
func VerifyCaptchaId(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"captcha_id": []string{"required"},
		"answer":     []string{"required"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"captcha_id": []string{
			"required:必填项不能为空",
		},
		"answer": []string{
			"required:必填项不能为空",
		},
	}
	return validate(data, rules, messages)
}

// 短信验证码
func VerifySmsVerifyCode(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}
