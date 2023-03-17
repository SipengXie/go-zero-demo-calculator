package svc

import (
	"go-zero-demo-calculator/calculator/calc/api/internal/config"
	"go-zero-demo-calculator/calculator/adder/rpc/adderclient"
	
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	AdderRpc adderclient.Adder
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AdderRpc : adderclient.NewAdder(zrpc.MustNewClient(c.AdderRpc)),
	}
}
