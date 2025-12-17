// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ai

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 聊天（流式）
func NewChatStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatStreamLogic {
	return &ChatStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatStreamLogic) ChatStream(req *types.ChatRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
