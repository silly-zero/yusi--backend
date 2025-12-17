// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/logic/ai"
	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
)

// AI 聊天（流式）
func ChatStreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ai.NewChatStreamLogic(r.Context(), svcCtx)
		resp, err := l.ChatStream(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
