package main

import (
	"flag"
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/model"
	"github.com/darrenli6/blog-server/internal/routers"
	"github.com/darrenli6/blog-server/pkg/logger"
	setting2 "github.com/darrenli6/blog-server/pkg/setting"
	"github.com/darrenli6/blog-server/pkg/tracer"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port    string
	runMode string
	config  string
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init setupDBengine err %v", err)
	}

	err = setupLogger()

	if err != nil {
		log.Fatalf("init logger err %v", err)
	}

	err = setFlag()
	if err != nil {
		log.Fatalf("init logger err %v", err)
	}
}

func setFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定配置文件路径")
	flag.Parse()

	return nil
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTrace(
		"block-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
func setupSetting() error {

	setting, err := setting2.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
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

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

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

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
