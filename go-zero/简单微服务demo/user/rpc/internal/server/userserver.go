// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero_study/user/rpc/internal/logic"
	"go-zero_study/user/rpc/internal/svc"
	"go-zero_study/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
