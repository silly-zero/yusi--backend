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

// 获取日记列表
func GetDiaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DiaryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := diary.NewGetDiaryListLogic(r.Context(), svcCtx)
		resp, err := l.GetDiaryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
