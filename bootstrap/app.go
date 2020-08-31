package bootstrap

import (
	"github.com/jinzhu/gorm"
	"github.com/maogou/ginapi/app/model"
	"github.com/maogou/ginapi/pkg/setting"
	"time"
)

//全局变量
var (
	ServeSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine *gorm.DB
)

//初始化配置
func InitSetting() error  {
	setting,err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = setting.ReadSection("Server",&ServeSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App",&AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database",&DatabaseSetting)
	if err != nil {
		return err
	}

	ServeSetting.ReadTimeout *= time.Second
	ServeSetting.WriteTimeout *= time.Second

	return nil
}

//实例化db引擎
func InitDBEngine() error  {
	var err error

	DBEngine,err = model.NewDBEngine(DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}



