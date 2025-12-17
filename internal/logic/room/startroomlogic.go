// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始房间
func NewStartRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartRoomLogic {
	return &StartRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartRoomLogic) StartRoom(req *types.StartRoomRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
