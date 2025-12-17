// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package room

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取报告
func NewGetReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReportLogic {
	return &GetReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReportLogic) GetReport() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
