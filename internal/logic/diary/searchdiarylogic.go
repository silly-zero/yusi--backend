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

type SearchDiaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// 搜索日记
func NewSearchDiaryLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SearchDiaryLogic {
	return &SearchDiaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *SearchDiaryLogic) SearchDiary(keyword string, pageNum, pageSize int) (resp *types.Response, err error) {
	// 验证参数
	if keyword == "" {
		return &types.Response{
			Code:    400,
			Message: "搜索关键词不能为空",
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

	// 设置默认值
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 计算偏移量
	offset := (pageNum - 1) * pageSize

	// 搜索日记（在标题和内容中搜索）
	var diaries []model.Diary
	var total int64

	// 构建查询条件 - 只搜索自己的日记
	query := l.svcCtx.DB.Where("user_id = ?", userId).
		Where("(title LIKE ? OR content LIKE ?)", "%"+keyword+"%", "%"+keyword+"%")

	// 获取总数
	query.Model(&model.Diary{}).Count(&total)

	// 查询日记列表，按创建时间倒序
	result := query.Order("create_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&diaries)

	if result.Error != nil {
		return &types.Response{
			Code:    500,
			Message: "搜索失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data: map[string]interface{}{
			"total":   total,
			"list":    diaries,
			"page":    pageNum,
			"perPage": pageSize,
		},
	}, nil
}
