package v1

import (
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/service"
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/darrenli6/blog-server/pkg/convert"
	"github.com/darrenli6/blog-server/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

	param := service.TagListRequest{}

	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindingAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc CountTag %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFailed)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc TagList %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFailed)
		return
	}

	response.ToResponseList(tags, totalRows)
	return

}

// @Summary create标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/create [get]
func (t Tag) Create(c *gin.Context) {

	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindingAndValid errs: %v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err1 := svc.CreateTag(&param)
	if err1 != nil {
		global.Logger.Errorf("svc CreateTag %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFailed)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary update标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/update [get]
func (t Tag) Update(c *gin.Context) {

	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindingAndValid errs: %v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	err1 := svc.UpdateTag(&param)

	if err1 != nil {
		global.Logger.Errorf("svc UpdateTag %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFailed)
		return
	}

	response.ToResponse(gin.H{})
	return

}

// @Summary delete标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Router /api/v1/tags/delete [get]
func (t Tag) Delete(c *gin.Context) {

	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}

	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindingAndValid errs: %v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())

	err1 := svc.DeleteTag(&param)

	if err1 != nil {
		global.Logger.Errorf("svc DeleteTag %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFailed)
		return
	}

	response.ToResponse(gin.H{})
	return

}
