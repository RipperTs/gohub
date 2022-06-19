// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string
	Email    string
	Phone    string `json:"-"` // 忽略字段输出
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (this *User) Create() {
	database.DB.Create(&this)
}

// ComparePassword 密码是否正确
func (this *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, this.Password)
}

// Save 保存数据
func (this *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&this)
	return result.RowsAffected
}
