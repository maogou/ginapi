package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/bootstrap"
	"github.com/maogou/ginapi/routes"
	"log"
	"net/http"
	"time"
)

//初始化方法-全局变量-init方法-main方法执行流程
func init()  {
	err := bootstrap.InitSetting()
	if err != nil {
		log.Fatalf("bootstrap.InitSetting err: %v",err)
	}

	err = bootstrap.InitLogger()
	if err != nil {
		log.Fatalf("bootstrap.InitLogger: %v",err)
	}
}

//应用主入口
func main() {
	//设置运行模式
	gin.SetMode(bootstrap.ServeSetting.RunMode)

	httpPort := ":" + bootstrap.ServeSetting.HttpPort
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


