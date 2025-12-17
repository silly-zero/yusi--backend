// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"yusi-backend/internal/svc"
	"yusi-backend/internal/types"
	"yusi-backend/internal/utils"
	"yusi-backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.Response, err error) {
	// 验证参数
	if req.UserName == "" || req.Password == "" {
		return &types.Response{
			Code:    400,
			Message: "用户名和密码不能为空",
		}, nil
	}

	// 检查用户名是否已存在
	var existingUser model.User
	result := l.svcCtx.DB.Where("user_name = ?", req.UserName).First(&existingUser)
	if result.Error == nil {
		return &types.Response{
			Code:    400,
			Message: "用户名已存在",
		}, nil
	}

	// 检查邮箱是否已存在（如果提供了邮箱）
	if req.Email != "" {
		result = l.svcCtx.DB.Where("email = ?", req.Email).First(&existingUser)
		if result.Error == nil {
			return &types.Response{
				Code:    400,
				Message: "邮箱已被使用",
			}, nil
		}
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "密码加密失败",
		}, nil
	}

	// 创建用户
	user := model.User{
		UserId:   utils.GenerateID(),
		UserName: req.UserName,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := l.svcCtx.DB.Create(&user).Error; err != nil {
		return &types.Response{
			Code:    500,
			Message: "创建用户失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "注册成功",
		Data: map[string]interface{}{
			"userId":   user.UserId,
			"userName": user.UserName,
			"email":    user.Email,
		},
	}, nil
}
