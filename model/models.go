package model

import (
	"time"
)

// User 用户模型
type User struct {
	UserId     string    `gorm:"column:user_id;primaryKey" json:"userId"`
	UserName   string    `gorm:"column:user_name" json:"userName"`
	Password   string    `gorm:"column:password" json:"-"`
	Email      string    `gorm:"column:email" json:"email"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

// Diary 日记模型
type Diary struct {
	DiaryId    string    `gorm:"column:diary_id;primaryKey" json:"diaryId"`
	UserId     string    `gorm:"column:user_id;index" json:"userId"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content;type:text" json:"content"`
	Visibility bool      `gorm:"column:visibility" json:"visibility"`
	EntryDate  time.Time `gorm:"column:entry_date" json:"entryDate"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (Diary) TableName() string {
	return "diary"
}
