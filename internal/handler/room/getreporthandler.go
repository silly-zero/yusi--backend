// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/room"
	"yusi-backend/internal/svc"
)

// 获取报告
func GetReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := room.NewGetReportLogic(r.Context(), svcCtx)
		resp, err := l.GetReport()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
