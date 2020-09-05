package dao

import "github.com/jinzhu/gorm"

//对数据层封装基类
type Dao struct {
	engine *gorm.DB
}

//实例化dao
func New(engine *gorm.DB) *Dao  {
	return &Dao{
		engine: engine,
	}
}
