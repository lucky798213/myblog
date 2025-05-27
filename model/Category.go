package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null " json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("username = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.Error_CATENAME_USED //1001
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
