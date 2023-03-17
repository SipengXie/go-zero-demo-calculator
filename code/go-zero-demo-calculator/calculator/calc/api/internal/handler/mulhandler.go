package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo-calculator/calculator/calc/api/internal/logic"
	"go-zero-demo-calculator/calculator/calc/api/internal/svc"
	"go-zero-demo-calculator/calculator/calc/api/internal/types"
)

func mulHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MulReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMulLogic(r.Context(), svcCtx)
		resp, err := l.Mul(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
