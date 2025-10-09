package model

import "time"

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	Nickname  string
	AvatarUrl string
	Email     string `gorm:"unique"`
	Wechat    string `gorm:"unique"`
	Bio       string
	AboutMe   string
	AboutBlog string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	SocialAccounts []SocialAccount `gorm:"foreignKey:UserID"`
}
