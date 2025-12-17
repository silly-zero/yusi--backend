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

type EditDiaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑日记
func NewEditDiaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditDiaryLogic {
	return &EditDiaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditDiaryLogic) EditDiary(req *types.EditDiaryRequest) (resp *types.Response, err error) {
	// 验证参数
	if req.DiaryId == "" {
		return &types.Response{
			Code:    400,
			Message: "日记ID不能为空",
		}, nil
	}

	// 查询日记是否存在
	var diary model.Diary
	result := l.svcCtx.DB.Where("diary_id = ?", req.DiaryId).First(&diary)
	if result.Error != nil {
		return &types.Response{
			Code:    404,
			Message: "日记不存在",
		}, nil
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	// Visibility 是bool类型，需要特殊处理
	updates["visibility"] = req.Visibility

	if err := l.svcCtx.DB.Model(&diary).Updates(updates).Error; err != nil {
		return &types.Response{
			Code:    500,
			Message: "更新日记失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "更新成功",
	}, nil
}
