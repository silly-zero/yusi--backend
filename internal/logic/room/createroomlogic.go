// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建房间
func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreateRoomRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
