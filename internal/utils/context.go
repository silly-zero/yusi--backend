package utils

import (
	"context"
	"errors"
	"net/http"
)

type contextKey string

const (
	UserIdKey   contextKey = "userId"
	UserNameKey contextKey = "userName"
)

// SetUserId 设置用户ID到上下文
func SetUserId(r *http.Request, userId string) *http.Request {
	ctx := context.WithValue(r.Context(), UserIdKey, userId)
	return r.WithContext(ctx)
}

// SetUserName 设置用户名到上下文
func SetUserName(r *http.Request, userName string) *http.Request {
	ctx := context.WithValue(r.Context(), UserNameKey, userName)
	return r.WithContext(ctx)
}

// GetUserId 从上下文获取用户ID
func GetUserId(r *http.Request) (string, error) {
	userId, ok := r.Context().Value(UserIdKey).(string)
	if !ok || userId == "" {
		return "", errors.New("未找到用户ID")
	}
	return userId, nil
}

// GetUserName 从上下文获取用户名
func GetUserName(r *http.Request) (string, error) {
	userName, ok := r.Context().Value(UserNameKey).(string)
	if !ok || userName == "" {
		return "", errors.New("未找到用户名")
	}
	return userName, nil
}
