package user

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (this User) {
	database.DB.Where("phone = ?", phone).First(&this)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (this User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&this)
	return
}

// Get 通过 ID 获取用户
func Get(idstr string) (this User) {
	database.DB.Where("id", idstr).First(&this)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (this User) {
	database.DB.Where("email = ?", email).First(&this)
	return
}

// All 获取所有数据
func All() (this []User) {
	database.DB.Find(&this)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (this []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&this,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
