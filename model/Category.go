package model

import "Q-BLOG/utils/errcode"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCate 查询分类是否存在
func CheckCate(username string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", username).First(&cate)
	if cate.ID > 0 {
		return errcode.ERROR_CATENAME_USED //1001
	}
	return errcode.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errcode.ERROR //500
	}
	return errcode.SUCCESS //200
}

// 查询分类下的所有文章

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil {
		return nil
	}
	return cate
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err2 := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err2 != nil {
		return errcode.ERROR
	}
	return errcode.SUCCESS
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err3 := db.Where("id = ?", id).Delete(&cate).Error
	if err3 != nil {
		return errcode.ERROR
	}
	return errcode.SUCCESS
}
