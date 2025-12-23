// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/diary"
	"yusi-backend/internal/svc"
)

// 删除日记
func DeleteDiaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从URL路径获取diaryId
		diaryId := r.URL.Query().Get(":diaryId")
		if diaryId == "" {
			// 尝试从URL路径获取
			diaryId = r.URL.Path[len("/api/diary/"):]
		}

		l := diary.NewDeleteDiaryLogic(r.Context(), svcCtx, r)
		resp, err := l.DeleteDiary(diaryId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
