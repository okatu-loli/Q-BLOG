package model

import (
	"Q-BLOG/utils/errcode"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username""`
	Password string `gorm:"type:varchar(20);not null" json:"password""`
	Role     int    `gorm:"type:int" json:"role""`
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errcode.ERROR_USERNAME //1001
	}
	return errcode.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errcode.ERROR //500
	}
	return errcode.SUCCESS //200
}
