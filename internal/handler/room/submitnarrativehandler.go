// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/room"
	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
)

// 提交叙述
func SubmitNarrativeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitNarrativeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := room.NewSubmitNarrativeLogic(r.Context(), svcCtx)
		resp, err := l.SubmitNarrative(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
