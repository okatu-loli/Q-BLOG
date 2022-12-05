package model

import (
	"Q-BLOG/utils/errcode"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errcode.ERROR_USERNAME //1001
	}
	return errcode.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errcode.ERROR //500
	}
	return errcode.SUCCESS //200
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err2 := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err2 != nil {
		return errcode.ERROR
	}
	return errcode.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err3 := db.Where("id = ?", id).Delete(&user).Error
	if err3 != nil {
		return errcode.ERROR
	}
	return errcode.SUCCESS
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}
