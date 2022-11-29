package api

import (
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/service"
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/darrenli6/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {

	param := service.AuthRequest{}

	global.Logger.Infof(c, "auth")
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.binding error %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("svc.GenerateToken err %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
	return
}
