package svc

import (
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/admin/internal/config"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/models"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  userclient.User
	AuthUser *userclient.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: userclient.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
