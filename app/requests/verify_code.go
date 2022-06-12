package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodeRequest struct {
	Captcha_id string `json:"captcha_id,omitempty" valid:"captcha_id"`
	Answer     string `json:"answer,omitempty" valid:"answer"`
}

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
