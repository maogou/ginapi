package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/app/model"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/logger"
	"github.com/maogou/ginapi/pkg/setting"
	"github.com/maogou/ginapi/routes"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

//初始化方法-全局变量-init方法-main方法执行流程
func init() {
	//初始化配置
	err := initSetting()
	if err !=nil {
		log.Fatalf("init.initSetting err: %v",err)
	}

	//初始化日志
	err = initLogger()
	if err != nil {
		log.Fatalf("init.initLogger err: %v",err)
	}

	//初始化db
	err = initDBEngine()
	if err != nil {
		log.Fatalf("init.initDBEngine err: %v",err)
	}
}

//应用主入口
func main() {
	//设置运行模式
	gin.SetMode(global.ServeSetting.RunMode)

	httpPort := ":" + global.ServeSetting.HttpPort
	//实例化路由
	router := routes.NewRouter()

	//自定义serve
	serve := &http.Server{
		Addr: httpPort,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//运行服务
	serve.ListenAndServe()
}

//初始化配置
func initSetting() error  {
	setting,err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = setting.ReadSection("Server",&global.ServeSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServeSetting.ReadTimeout *= time.Second
	global.ServeSetting.WriteTimeout *= time.Second

	return nil
}

//实例化db引擎
func initDBEngine() error  {
	var err error

	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

func initLogger() error  {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, //600M
		MaxAge:    10, //10天
		LocalTime: true, //使用本地时间格式
	},"",log.LstdFlags).WithCaller(2)

	return nil
}


