package models

import (
	"time"

	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Summary     string         `json:"summary"`
	Status      string         `json:"status" gorm:"default:'draft'"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 外键
	UserID uint `json:"user_id" gorm:"not null"`
	
	// 关联关系
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
} 