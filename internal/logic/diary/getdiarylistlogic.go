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

type GetDiaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取日记列表
func NewGetDiaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiaryListLogic {
	return &GetDiaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiaryListLogic) GetDiaryList(req *types.DiaryListRequest) (resp *types.Response, err error) {
	// 验证参数
	if req.UserId == "" {
		return &types.Response{
			Code:    400,
			Message: "用户ID不能为空",
		}, nil
	}

	// 默认值
	if req.PageNum < 1 {
		req.PageNum = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 10
	}

	// 计算偏移量
	offset := (req.PageNum - 1) * req.PageSize

	// 查询总数
	var total int64
	l.svcCtx.DB.Model(&model.Diary{}).Where("user_id = ?", req.UserId).Count(&total)

	// 查询列表
	var diaries []model.Diary
	query := l.svcCtx.DB.Where("user_id = ?", req.UserId)

	// 排序
	if req.SortBy != "" {
		order := req.SortBy
		if !req.Asc {
			order += " DESC"
		} else {
			order += " ASC"
		}
		query = query.Order(order)
	} else {
		// 默认按创建时间倒序
		query = query.Order("create_time DESC")
	}

	result := query.Offset(offset).Limit(req.PageSize).Find(&diaries)
	if result.Error != nil {
		return &types.Response{
			Code:    500,
			Message: "查询日记列表失败",
		}, nil
	}

	// 构造响应
	listResp := types.DiaryListResponse{
		Total:   total,
		List:    make([]types.Diary, 0, len(diaries)),
		Page:    req.PageNum,
		PerPage: req.PageSize,
	}

	// 转换数据
	for _, d := range diaries {
		listResp.List = append(listResp.List, types.Diary{
			DiaryId:    d.DiaryId,
			UserId:     d.UserId,
			Title:      d.Title,
			Content:    d.Content,
			Visibility: d.Visibility,
			EntryDate:  d.EntryDate.Format("2006-01-02"),
			CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.Response{
		Code:    200,
		Message: "success",
		Data:    listResp,
	}, nil
}
