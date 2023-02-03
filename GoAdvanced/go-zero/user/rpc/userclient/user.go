// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserAuthReply   = user.UserAuthReply
	UserAuthRequest = user.UserAuthRequest

	User interface {
		Auth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Auth(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserAuthReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Auth(ctx, in, opts...)
}
