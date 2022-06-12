package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (this *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(this.Password) {
		this.Password = hash.BcryptHash(this.Password)
	}
	return
}
