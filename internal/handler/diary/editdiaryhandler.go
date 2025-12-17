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

// 编辑日记
func EditDiaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditDiaryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := diary.NewEditDiaryLogic(r.Context(), svcCtx)
		resp, err := l.EditDiary(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
