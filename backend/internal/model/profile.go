package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID      uint      `gorm:"uniqueIndex;not null" json:"userID"`
	AvatarURL   string    `gorm:"size:255" json:"avatarUrl"`
	Nickname    string    `gorm:"size:50" json:"nickname"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Email       string    `gorm:"size:100" json:"email"`
	Wechat      string    `gorm:"size:50" json:"wechat"`
	AboutMe     string    `gorm:"type:text" json:"aboutMe"`
	AboutBlog   string    `gorm:"type:text" json:"aboutBlog"`
	SocialMedia string    `gorm:"type:json" json:"socialMedia"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type SocialAccount struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
