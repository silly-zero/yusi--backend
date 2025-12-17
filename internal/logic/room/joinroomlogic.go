// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加入房间
func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinRoomLogic) JoinRoom(req *types.JoinRoomRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
