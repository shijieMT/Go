package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero_study/user/rpc/userclient"
	"go-zero_study/video/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// 新增 UserRpc userclient.User
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 新增 UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
