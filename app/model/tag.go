package model

import (
	"github.com/jinzhu/gorm"
	"github.com/maogou/ginapi/pkg/app"
)

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}


//使用gorm库手动指定表名
func (t Tag)TableName() string{
	return "blog_tag"
}

//针对swagger生成文档包含别的结构体无法显示的bug
type TagSwagger struct {
	List []*Tag
	Pager *app.Page
}

//查询tag的总条数
func (t Tag)Count(db *gorm.DB) (int,error)  {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}

	db = db.Where("state = ?",t.State)

	if err := db.Model(&t).Where("is_del = ?",0).Count(&count).Error;err != nil {
		return 0,err
	}

	return count,nil
}

//获取tag的列表
func (t Tag)List(db *gorm.DB,pageOffset,pageSize int)([]*Tag,error)  {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}

	db = db.Where("state = ?",t.State)

	if err = db.Where("is_del = ?",0).Find(&tags).Error;err != nil {
		return nil,err
	}

	return tags,nil
}

//新增tag标签
func (t Tag)Create(db *gorm.DB) error  {
	return db.Create(&t).Error
}

//更新标签
func (t Tag)Update(db *gorm.DB,values interface{})error  {
	return db.Model(&t).Where("id = ? AND is_del = ?",t.ID,0).Updates(&t).Error
}

//删除标签
func (t Tag)Delete(db *gorm.DB) error  {
	return db.Where("id = AND is_del = ?",t.Model.ID,0).Delete(&t).Error
}
