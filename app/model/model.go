package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maogou/ginapi/bootstrap"
	"github.com/maogou/ginapi/pkg/setting"
	"time"
)

//所有model的公共属性
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))

	if err != nil {
		return nil, err
	}

	//开启sql日志
	if bootstrap.ServeSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetConnMaxIdleTime(time.Duration(databaseSetting.MaxIdleConns))
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
