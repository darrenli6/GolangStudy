package main

import (
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/model"
	"github.com/darrenli6/blog-server/internal/routers"
	"github.com/darrenli6/blog-server/pkg/logger"
	setting2 "github.com/darrenli6/blog-server/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupsetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBengine err %v", err)
	}

	err = setupLogger()

	if err != nil {
		log.Fatalf("init logger err %v", err)
	}
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,

		// 600M
		MaxSize: 600,
		// 生存周期10天
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil

}

// @title 博客系统
// @version 1.0
// @description go的项目
// @termsOfService https://github.com/darrenli/
func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	global.Logger.Infof("%s ", "main")
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
