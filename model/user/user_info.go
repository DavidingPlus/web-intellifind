package user

import (
	"backend/tools/hash"
	"backend/utils"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type UserInfos struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email" valid:"email"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avator" gorm:"column:avator"` //头像
	City     string `json:"city" gorm:"column:city"`     //城市

}

// 获取用户列表
func GetUserList() []*UserInfos {
	data := make([]*UserInfos, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func GetUser(key string) UserInfos {
	user := UserInfos{}
	utils.DB.Where("email = ?", key).Or("id= ?", key).First(&user)
	return user
}

// 新建用户
func CreateUser(user UserInfos) {
	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	utils.DB.Create(&user)
}

// 删除用户
func DelteUser(id_string string) {
	id, _ := strconv.Atoi(id_string)
	utils.DB.Where("id = ?", id).Delete(&UserInfos{})
}

func GetPasswoord(key string) string {
	user := UserInfos{}
	utils.DB.Where("email = ?", key).Or("id = ?", key).First(&user)
	return user.Password
}

func ComparePassword(_password string, email string) bool {
	real_password := GetPasswoord(email)
	return hash.BcryptCheck(_password, real_password)
}

// 更新用户密码
func UpdatePassword(password string, key string) {
	user := UserInfos{}
	utils.DB.Where("email = ?", key).Or("id = ?", key).First(&user)
	user.Password = hash.BcryptHash(password)
	utils.DB.Save(&user)
}

// 保存
func (userModel *UserInfos) Save() (rowsAffected int64) {
	result := utils.DB.Save(&userModel)
	return result.RowsAffected
}
