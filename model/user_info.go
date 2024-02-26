package model

import (
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

// 通过id获取个人信息
func GetUserInfoById(id uint) UserInfos {
	user := UserInfos{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}

// 通过名字获取个人信息
func GetUserByUsername(username string) UserInfos {
	user := UserInfos{}
	utils.DB.Where("username = ?", username).First(&user)
	return user
}

// 通过email获取个人信息
func GetUserByEmail(email string) UserInfos {
	user := UserInfos{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

// 新建用户
func CreateUser(user UserInfos) {
	utils.DB.Create(&user)
}

// 删除用户
func DelteUser(id_string string) {
	id, _ := strconv.Atoi(id_string)
	utils.DB.Where("id = ?", id).Delete(&UserInfos{})
}
