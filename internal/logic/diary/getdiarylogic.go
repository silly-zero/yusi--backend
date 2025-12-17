// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
	"yusi-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取日记详情
func NewGetDiaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiaryLogic {
	return &GetDiaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiaryLogic) GetDiary(diaryId string) (resp *types.Response, err error) {
	// 验证参数
	if diaryId == "" {
		return &types.Response{
			Code:    400,
			Message: "日记ID不能为空",
		}, nil
	}

	// 查询日记
	var diary model.Diary
	result := l.svcCtx.DB.Where("diary_id = ?", diaryId).First(&diary)
	if result.Error != nil {
		return &types.Response{
			Code:    404,
			Message: "日记不存在",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    diary,
	}, nil
}
