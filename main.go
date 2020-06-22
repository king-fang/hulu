package main

import (
	"context"
	"fmt"
	"hulujia/config"
	"hulujia/conn"
	"hulujia/form"
	"hulujia/router"
	"hulujia/util/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title 葫芦家
// @version 2.0
// @description 葫芦家商城v2
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @contact.name flaravel
// @host localhost:8100
func main()  {

	// 1. 初始化配置
	config.SetupConfig()

	// 2.初始化路由
	routers := router.SetupRouter()

	// 3.连接mysql数据库
	conn.SetupMysql()

	// 4.初始化表单验证
	form.SetUp()

	// 5.开启服务
	server := &http.Server{
		Addr         : fmt.Sprintf(":%d",config.App.Port),
		Handler      : routers,
		ReadTimeout  : config.App.AppReadTimeout,
		WriteTimeout : config.App.AppWriteTimeout,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Go hTTP server listen: %s\n", err)
		}
	}()
	log.Info("Go http server Port:" + fmt.Sprintf("%d", config.App.Port) + "\tPid:" + fmt.Sprintf("%d", os.Getpid()))
	log.Info("Go http server start successful")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Info("Get Signal:", sig)
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
