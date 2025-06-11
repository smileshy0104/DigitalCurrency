package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/uc_client"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	URegisterRpc uc_client.Register
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		URegisterRpc: uc_client.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
