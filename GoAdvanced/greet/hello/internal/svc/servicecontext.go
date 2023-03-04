package svc

import (
	"github.com/darrenli6/GolangStudy/GoAdvanced/greet/hello/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
