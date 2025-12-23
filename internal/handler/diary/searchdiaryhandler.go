// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/diary"
	"yusi-backend/internal/svc"
)

// 搜索日记
func SearchDiaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从查询参数获取
		keyword := r.URL.Query().Get("keyword")
		pageNumStr := r.URL.Query().Get("pageNum")
		pageSizeStr := r.URL.Query().Get("pageSize")

		// 解析分页参数
		pageNum, _ := strconv.Atoi(pageNumStr)
		pageSize, _ := strconv.Atoi(pageSizeStr)

		if pageNum < 1 {
			pageNum = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}

		l := diary.NewSearchDiaryLogic(r.Context(), svcCtx, r)
		resp, err := l.SearchDiary(keyword, pageNum, pageSize)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
