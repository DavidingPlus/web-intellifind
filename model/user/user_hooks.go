package user

import (
	"backend/tools/hash"
	"gorm.io/gorm"
)

func (u *UserInfos) BeforeCreate(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(u.Password) {
		u.Password = hash.BcryptHash(u.Password)
	}
	return
}
