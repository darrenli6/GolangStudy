package logic

import (
	"context"
	"errors"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/helper"

	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/rpc/internal/svc"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	// todo: add your logic here and delete this line
	if in.Token == "" {
		return nil, errors.New("必填参数不能为空")
	}
	claim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Identity = claim.Identity
	resp.Id = uint64(claim.Id)
	resp.Extend = map[string]string{
		"name": claim.Name,
	}
	return &user.UserAuthReply{}, nil
}
