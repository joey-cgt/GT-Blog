package repository

import (
	"context"
	"gt-blog/backend_ddd/internal/admin_context/model"

	"gorm.io/gorm"
)

// AdminRepository 定义了与 Admin 实体相关的所有数据操作接口
// 这样做的好处是：
// 1. 清晰地定义了数据访问层的职责。
// 2. 便于后续进行单元测试（可以轻松地用 mock 实现来替代真实的数据库操作）。
type AdminRepository interface {
	GetAdminByID(ctx context.Context, id uint) (*model.Admin, error)
	GetAdminByUsername(ctx context.Context, username string) (*model.Admin, error)
	UpdateAdmin(ctx context.Context, admin *model.Admin) error
}

type MySqlAdminRepository struct {
	db *gorm.DB
}

func NewMySqlAdminRepository(db *gorm.DB) *MySqlAdminRepository {
	return &MySqlAdminRepository{
		db: db,
	}
}

// GetAdminByID 根据ID获取管理员信息
func (m *MySqlAdminRepository) GetAdminByID(ctx context.Context, id uint) (*model.Admin, error) {
	var admin model.Admin
	if err := m.db.WithContext(ctx).Preload("SocialAccounts").Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (m *MySqlAdminRepository) GetAdminByUsername(ctx context.Context, username string) (*model.Admin, error) {
	var admin model.Admin
	if err := m.db.WithContext(ctx).Preload("SocialAccounts").Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (m *MySqlAdminRepository) UpdateAdmin(ctx context.Context, admin *model.Admin) error {
	// 先保存admin基本信息
	if err := m.db.WithContext(ctx).Save(admin).Error; err != nil {
		return err
	}

	// 处理社交账号关联
	if len(admin.SocialAccounts) > 0 {
		// 删除旧的社交账号
		if err := m.db.WithContext(ctx).Where("user_id = ?", admin.ID).Delete(&model.SocialAccount{}).Error; err != nil {
			return err
		}

		// 更新非空的社交账号
		var validSocialAccounts []model.SocialAccount
		for _, account := range admin.SocialAccounts {
			if account.Platform != "" && account.Url != "" {
				// 确保设置正确的UserID
				account.UserID = admin.ID
				validSocialAccounts = append(validSocialAccounts, account)
			}
		}

		// 如果有有效的社交账号，则批量创建
		if len(validSocialAccounts) > 0 {
			if err := m.db.WithContext(ctx).Create(&validSocialAccounts).Error; err != nil {
				return err
			}
		}

	}

	return nil
}
