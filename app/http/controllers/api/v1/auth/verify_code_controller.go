package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/sms"
	"math/rand"
	"strconv"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)

	// 返回响应结果
	response.ShowSuccess(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// VerifyCaptcha 验证图片验证码
func (vc *VerifyCodeController) VerifyCaptcha(c *gin.Context) {
	// 表单验证
	request := requests.VerifyCodeRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCaptchaId); !ok {
		return
	}

	// 验证码是否正确
	result := captcha.NewCaptcha().VerifyCaptcha(request.Captcha_id, request.Answer)
	// 返回响应结果
	response.ShowSuccess(c, result)
}

// SendSmsVerifyCode 发送短信验证码
func (vc *VerifyCodeController) SendSmsVerifyCode(c *gin.Context) {
	// 表单验证
	request := requests.SmsVerifyCodeRequest{}
	if ok := requests.Validate(c, &request, requests.VerifySmsVerifyCode); !ok {
		return
	}

	// 随机生成6位数字
	code := rand.Intn(999999) + 100000
	// 执行发送短信验证码操作
	result := sms.NewSMS().Send(request.Phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": strconv.Itoa(code)},
	}, 1)

	if result {
		response.ShowSuccess(c, result, "短信验证码发送成功!")
		return
	}
	response.ShowError(c, 422, "短信验证码发送失败!")
}
