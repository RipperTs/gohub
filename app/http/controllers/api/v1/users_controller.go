package v1

import (
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.ShowSuccess(c, userModel)
}

// Index 查询Users列表(分页)
func (this *UsersController) Index(c *gin.Context) {
	// 验证请求参数
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pagination := user.Paginate(c, 30)
	response.ShowSuccess(c, gin.H{
		"list":       data,
		"pagination": pagination,
	})
}
