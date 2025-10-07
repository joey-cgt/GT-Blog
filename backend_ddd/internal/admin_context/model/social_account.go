package model

import "time"

type SocialAccount struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Platform  string
	Url       string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
