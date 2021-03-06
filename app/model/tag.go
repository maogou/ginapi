package model

import (
	"github.com/maogou/ginapi/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

//使用gorm库手动指定表名
func (t Tag) TableName() string {
	return "blog_tag"
}

//针对swagger生成文档包含别的结构体无法显示的bug
type TagSwagger struct {
	List  []*Tag
	Pager *app.Page
}

//查询tag的总条数
func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//获取tag的列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

//新增tag标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

//更新标签
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Updates(values).Where("id = ? AND is_del = ?", t.ID, 0).Error
}

//删除标签
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?  AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

func (t Tag) ListByIDs(db *gorm.DB, ids []uint32) ([]*Tag, error) {
	var tags []*Tag
	db = db.Where("state = ? AND is_del = ?", t.State, 0)
	err := db.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}
