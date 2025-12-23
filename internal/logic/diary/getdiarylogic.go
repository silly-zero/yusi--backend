// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package diary

import (
	"context"
	"net/http"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
	"yusi-backend/internal/utils"
	"yusi-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// 获取日记详情
func NewGetDiaryLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetDiaryLogic {
	return &GetDiaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
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

	// 获取当前用户ID
	userId, err := utils.GetUserId(l.r)
	if err != nil || userId == "" {
		return &types.Response{
			Code:    401,
			Message: "未授权",
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

	// 验证权限：只能查看自己的日记
	if diary.UserId != userId {
		return &types.Response{
			Code:    403,
			Message: "无权限查看此日记",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    diary,
	}, nil
}
