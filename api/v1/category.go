package v1

import (
	"Q-BLOG/model"
	"Q-BLOG/utils/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCate 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	msg = model.CheckCate(data.Name)
	if msg == errcode.SUCCESS {
		model.CreateCate(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"data":   data,
		"msg":    errcode.GetErrMsg(msg),
	})
}

// todo 查询分类下的所有文章

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page-size"))
	pageNum, _ := strconv.Atoi(c.Query("page-num"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data := model.GetCate(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"data":   data,
		"msg":    errcode.GetErrMsg(msg),
	})
}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	msg = model.CheckCate(data.Name)
	if msg == errcode.SUCCESS {
		model.EditCate(id, &data)
	}
	if msg == errcode.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"msg":    errcode.GetErrMsg(msg),
	})
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	msg = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status": msg,
		"msg":    errcode.GetErrMsg(msg),
	})
}
