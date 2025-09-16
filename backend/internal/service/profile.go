package service

import (
	"gt-blog/backend/internal/model"
	"gt-blog/backend/internal/repository"
)

func GetProfileByID(id uint) (*model.Profile, error) {
	// 简单实现，实际项目应添加缓存等逻辑
	var profile model.Profile
	err := repository.DB.Where("id = ?", id).First(&profile).Error
	return &profile, err
}