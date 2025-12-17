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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	// 验证参数
	if req.UserName == "" || req.Password == "" {
		return &types.Response{
			Code:    400,
			Message: "用户名和密码不能为空",
		}, nil
	}

	// 查询用户
	var user model.User
	result := l.svcCtx.DB.Where("user_name = ?", req.UserName).First(&user)
	if result.Error != nil {
		return &types.Response{
			Code:    401,
			Message: "用户名或密码错误",
		}, nil
	}

	// 验证密码
	if !utils.CheckPassword(user.Password, req.Password) {
		return &types.Response{
			Code:    401,
			Message: "用户名或密码错误",
		}, nil
	}

	// 生成 JWT Token
	token, err := utils.GenerateToken(
		user.UserId,
		user.UserName,
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
	)
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "生成token失败",
		}, nil
	}

	return &types.Response{
		Code:    200,
		Message: "登录成功",
		Data: map[string]interface{}{
			"token":    token,
			"userId":   user.UserId,
			"userName": user.UserName,
			"email":    user.Email,
		},
	}, nil
}
