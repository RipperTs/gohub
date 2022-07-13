package app

import (
	"github.com/idoubi/goz"
	"github.com/spf13/cast"
	"gohub/pkg/logger"
)

//get请求
func CurlGet(url string) goz.ResponseBody {
	cli := goz.NewClient()
	resp, err := cli.Get(url)
	if err != nil {
		logger.Error(cast.ToString(err))
	}
	content, err := resp.GetBody()
	if err != nil {
		logger.Error(cast.ToString(err))
	}
	return content
}

//post请求
func CurlPost(url string, Headers map[string]interface{}, FormParams map[string]interface{}) goz.ResponseBody {
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Headers:    Headers,
		FormParams: FormParams,
	})
	if err != nil {
		logger.Error(cast.ToString(err))
	}
	content, err := resp.GetBody()
	if err != nil {
		logger.Error(cast.ToString(err))
	}
	return content
}
