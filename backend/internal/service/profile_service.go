package service

import (
	"fmt"
	"gt-blog/backend/internal/model"
	"gt-blog/backend/internal/repository"
	"mime/multipart"
	"path/filepath"
	"time"
)

func GetProfileByID(id uint) (*model.Profile, error) {
	return repository.GetProfileByID(id)
}

func UploadAvatar(userID uint, file *multipart.FileHeader) (string, error) {
	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 存储路径
	uploadPath := filepath.Join("backend", "uploads", "avatars", newFilename)

	// 保存文件
	if err := repository.SaveAvatarFile(file, uploadPath); err != nil {
		return "", fmt.Errorf("保存头像文件失败: %v", err)
	}

	// 更新用户头像URL
	avatarURL := "/uploads/avatars/" + newFilename
	if err := repository.UpdateProfile(userID, map[string]interface{}{
		"avatar_url": avatarURL,
	}); err != nil {
		return "", fmt.Errorf("更新头像URL失败: %v", err)
	}

	return avatarURL, nil
}

func UpdateProfileByID(id uint, updateData map[string]interface{}) error {
	// 可更新字段白名单
	// 字段白名单（使用数据库字段命名规范）
	// 数据库字段白名单
	// allowedFields := map[string]bool{
	// 	"user_id":      true,
	// 	"nickname":     true,
	// 	"avatar_url":   true,
	// 	"bio":          true,
	// 	"email":        true,
	// 	"wechat":       true,
	// 	"blog_about":   true,
	// 	"social_media": true,
	// }

	// // 过滤非法字段并转换字段名
	// filteredData := make(map[string]interface{})
	// for k, v := range updateData {
	// 	// 检查是否为数据库字段名
	// 	if allowedFields[k] {
	// 		filteredData[k] = v
	// 	}
	// }

	// if len(filteredData) == 0 {
	// 	return fmt.Errorf("没有有效的更新字段")
	// }

	return repository.UpdateProfile(id, updateData)
}
