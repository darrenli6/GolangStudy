package global

import (
	"github.com/darrenli6/blog-server/pkg/logger"
	"github.com/darrenli6/blog-server/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)
