package v1

import (
	"Q-BLOG/model"
	"Q-BLOG/utils/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var msg int

// 查询用户是否存在
func userExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg = model.CheckUser(data.Username)
	if msg == errcode.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"data":   data,
		"msg":    errcode.GetErrMsg(msg),
	})
}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"data":   data,
		"msg":    errcode.GetErrMsg(msg),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}
