package service

import (
	"context"
	"errors"
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/dao"
	"github.com/darrenli6/blog-server/pkg/upload"
	"mime/multipart"
	"os"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save Directory")
		}

	}

	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("excessded maximum file limit")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("file permissions")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil

}