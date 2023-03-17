package logic

import (
	"context"

	"go-zero-demo-calculator/calculator/calc/api/internal/svc"
	"go-zero-demo-calculator/calculator/calc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MulLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMulLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MulLogic {
	return &MulLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MulLogic) Mul(req *types.MulReq) (resp *types.MulReply, err error) {

	return &types.MulReply{
		Res: req.A * req.B,
	}, nil
}

