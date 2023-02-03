package handler

import (
	"net/http"

	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/internal/logic"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/internal/svc"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
