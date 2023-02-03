package svc

import "github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
