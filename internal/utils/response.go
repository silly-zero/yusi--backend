package utils

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yusi-backend/internal/types"
)

// Success 返回成功响应
func Success(w http.ResponseWriter, data interface{}) {
	resp := types.Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	httpx.OkJson(w, resp)
}

// Fail 返回失败响应
func Fail(w http.ResponseWriter, code int, message string) {
	resp := types.Response{
		Code:    code,
		Message: message,
	}
	httpx.WriteJson(w, http.StatusOK, resp)
}

// Error 返回错误响应
func Error(w http.ResponseWriter, err error) {
	resp := types.Response{
		Code:    500,
		Message: err.Error(),
	}
	httpx.WriteJson(w, http.StatusOK, resp)
}

// BadRequest 返回 400 错误
func BadRequest(w http.ResponseWriter, message string) {
	Fail(w, 400, message)
}

// Unauthorized 返回 401 错误
func Unauthorized(w http.ResponseWriter, message string) {
	Fail(w, 401, message)
}

// NotFound 返回 404 错误
func NotFound(w http.ResponseWriter, message string) {
	Fail(w, 404, message)
}
