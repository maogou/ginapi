package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting,error)  {
	vp := viper.New()
	//设置配置的名称
	vp.SetConfigName("app")
	//添加配置文件的目录,可以多次调用此方法
	vp.AddConfigPath("config/")
	//配置文件的类型
	vp.SetConfigType("json")

	//读取配置文件
	err := vp.ReadInConfig()

	if err != nil {
		return nil,err
	}

	return &Setting{
		vp: vp,
	},nil
}
