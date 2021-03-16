package global

import (
	"github.com/maogou/ginapi/pkg/logger"
	"github.com/maogou/ginapi/pkg/setting"
	"gorm.io/gorm"
)

//全局变量
var (
	ServeSetting    *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JwtSetting      *setting.JwtAuthSettings
)
