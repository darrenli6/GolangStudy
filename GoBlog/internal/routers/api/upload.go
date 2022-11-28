package api

import (
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/service"
	"github.com/darrenli6/blog-server/pkg/app"
	"github.com/darrenli6/blog-server/pkg/convert"
	"github.com/darrenli6/blog-server/pkg/errcode"
	"github.com/darrenli6/blog-server/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {

	response := app.NewResponse(c)

	file, fileHeader, err := c.Request.FormFile("file")

	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())

	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		global.Logger.Errorf("svc.upload err %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
	return

}
