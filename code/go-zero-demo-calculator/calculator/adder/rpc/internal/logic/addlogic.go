package logic

import (
	"context"

	"go-zero-demo-calculator/calculator/adder/rpc/internal/svc"
	"go-zero-demo-calculator/calculator/adder/rpc/types/adder"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *adder.AddRequest) (*adder.AddResponse, error) {

	return &adder.AddResponse{Res:in.A + in.B}, nil
}
