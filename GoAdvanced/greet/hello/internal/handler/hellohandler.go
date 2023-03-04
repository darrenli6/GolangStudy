package handler

import (
	"net/http"

	"github.com/darrenli6/GolangStudy/GoAdvanced/greet/hello/internal/logic"
	"github.com/darrenli6/GolangStudy/GoAdvanced/greet/hello/internal/svc"
	"github.com/darrenli6/GolangStudy/GoAdvanced/greet/hello/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
