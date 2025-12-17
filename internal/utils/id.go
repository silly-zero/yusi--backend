package utils

import (
	"github.com/google/uuid"
)

// GenerateID 生成唯一 ID
func GenerateID() string {
	return uuid.New().String()
}
