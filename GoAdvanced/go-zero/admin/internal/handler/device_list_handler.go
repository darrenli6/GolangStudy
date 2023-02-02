package handler

import (
	"net/http"

	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/admin/internal/logic"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/admin/internal/svc"
	"github.com/darrenli6/GolangStudy/GoAdvanced/go-zero/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeviceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeviceListLogic(r.Context(), svcCtx)
		resp, err := l.DeviceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
