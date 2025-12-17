// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/diary"
	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
)

// 写日记
func WriteDiaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WriteDiaryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := diary.NewWriteDiaryLogic(r.Context(), svcCtx)
		resp, err := l.WriteDiary(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
