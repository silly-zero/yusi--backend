// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.Response, err error) {
	// 对于 JWT 无状态认证，登出主要在客户端完成
	// 这里返回成功响应即可
	// 如果需要实现 token 黑名单，可以将 token 存入 Redis

	return &types.Response{
		Code:    200,
		Message: "登出成功",
	}, nil
}
