// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"context"
	"time"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
	"yusi-backend/internal/utils"
	"yusi-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type WriteDiaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 写日记
func NewWriteDiaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WriteDiaryLogic {
	return &WriteDiaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WriteDiaryLogic) WriteDiary(req *types.WriteDiaryRequest) (resp *types.Response, err error) {
	// 验证参数
	if req.Title == "" || req.Content == "" {
		return &types.Response{
			Code:    400,
			Message: "标题和内容不能为空",
		}, nil
	}

	// 解析时间
	var entryDate time.Time
	if req.EntryDate != "" {
		entryDate, err = time.Parse("2006-01-02", req.EntryDate)
		if err != nil {
			return &types.Response{
				Code:    400,
				Message: "日期格式错误，应为 YYYY-MM-DD",
			}, nil
		}
	} else {
		entryDate = time.Now()
	}

	// 创建日记
	diary := model.Diary{
		DiaryId:    utils.GenerateID(),
		UserId:     req.UserId,
		Title:      req.Title,
		Content:    req.Content,
		Visibility: req.Visibility,
		EntryDate:  entryDate,
	}

	if err := l.svcCtx.DB.Create(&diary).Error; err != nil {
		return &types.Response{
			Code:    500,
			Message: "创建日记失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "创建成功",
		Data: map[string]interface{}{
			"diaryId": diary.DiaryId,
		},
	}, nil
}
