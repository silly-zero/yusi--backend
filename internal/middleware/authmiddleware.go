// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"
	"strings"

	"yusi-backend/internal/utils"
)

type AuthMiddleware struct {
	Secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		Secret: secret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从 Header 获取 Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.Unauthorized(w, "缺少认证令牌")
			return
		}

		// 检查 Bearer 前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(w, "认证令牌格式错误")
			return
		}

		tokenString := parts[1]

		// 验证 Token
		claims, err := utils.ParseToken(tokenString, m.Secret)
		if err != nil {
			utils.Unauthorized(w, "认证令牌无效或已过期")
			return
		}

		// 将用户信息存入上下文
		r = utils.SetUserId(r, claims.UserId)
		r = utils.SetUserName(r, claims.UserName)

		// 传递给下一个处理器
		next(w, r)
	}
}
