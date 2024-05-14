package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 新增 UserRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
}
