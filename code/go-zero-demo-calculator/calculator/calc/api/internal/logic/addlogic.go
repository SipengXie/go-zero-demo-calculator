package logic

import (
	"context"

	"go-zero-demo-calculator/calculator/adder/rpc/types/adder"
	"go-zero-demo-calculator/calculator/calc/api/internal/svc"
	"go-zero-demo-calculator/calculator/calc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddReply, err error) {
	res, err := l.svcCtx.AdderRpc.Add(l.ctx, &adder.AddRequest{
		A: req.A,
		B: req.B,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddReply{
		Res: res.Res,
	}, nil
}

