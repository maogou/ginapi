package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/app/model"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/logger"
	"github.com/maogou/ginapi/pkg/setting"
	"github.com/maogou/ginapi/routes"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//初始化方法-全局变量-init方法-main方法执行流程
func init() {
	//初始化配置
	err := initSetting()
	if err != nil {
		log.Fatalf("init.initSetting err: %v", err)
	}

	//初始化日志
	err = initLogger()
	if err != nil {
		log.Fatalf("init.initLogger err: %v", err)
	}

	//初始化db
	err = initDBEngine()
	if err != nil {
		log.Fatalf("init.initDBEngine err: %v", err)
	}
}

//@title GinApi文档

//@version 1.0
//@description 使用gin框架开发api接口
//@termOfService https://juluzhizhan.com
func main() {
	//设置运行模式
	gin.SetMode(global.ServeSetting.RunMode)

	httpPort := ":" + global.ServeSetting.HttpPort
	//实例化路由
	router := routes.NewRouter()

	//自定义serve
	serve := &http.Server{
		Addr:           httpPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//优雅的重启
	go func() {
		//运行服务
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve.ListenAndServe err: %#v", err)
		}
	}()

	//等待信号中断
	quit := make(chan os.Signal)
	//接收syscall.SIGINT和syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown down server .....")

	//最大时间控制,通知该服务端它有5秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serve.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown:%#v", err)
	}

	log.Println("Server exiting")

}

//初始化配置
func initSetting() error {
	appSetting, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = appSetting.ReadSection("Server", &global.ServeSetting)
	if err != nil {
		return err
	}

	err = appSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = appSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = appSetting.ReadSection("JWT", &global.JwtSetting)

	if err != nil {
		return err
	}

	err = appSetting.ReadSection("Zap", &global.ZapLogSetting)
	if err != nil {
		return err
	}

	global.ServeSetting.ReadTimeout *= time.Second
	global.ServeSetting.WriteTimeout *= time.Second
	global.JwtSetting.Expire *= time.Second

	return nil
}

//实例化db引擎
func initDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

//初始化日志服务
func initLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,  //600M
		MaxAge:    10,   //10天
		LocalTime: true, //使用本地时间格式
	})

	return nil
}
