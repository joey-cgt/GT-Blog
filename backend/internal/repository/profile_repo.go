package repository

import (
	"gt-blog/backend/internal/model"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// GetProfileByID 根据ID获取用户资料
func GetProfileByID(id uint) (*model.Profile, error) {
	var profile model.Profile
	if err := DB.First(&profile, id).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

// SaveAvatarFile 保存头像文件到指定路径
func SaveAvatarFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标目录
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// 复制文件内容
	_, err = io.Copy(out, src)
	return err
}

// UpdateProfile 更新用户资料
func UpdateProfile(id uint, data map[string]interface{}) error {
	return DB.Model(&model.Profile{}).
		Where("id = ?", id).
		Updates(data).Error
}
