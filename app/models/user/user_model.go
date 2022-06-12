// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (m *User) Create() {
	database.DB.Create(&m)
}
