package main

import (
	"github.com/maogou/ginapi/routes"
	"net/http"
	"time"
)

func main() {

	router := routes.NewRouter()

	//自定义serve
	serve := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	serve.ListenAndServe()
}
