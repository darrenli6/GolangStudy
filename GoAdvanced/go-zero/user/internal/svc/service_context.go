package svc

import (
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/models"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
