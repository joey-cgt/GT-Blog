package repository

import (
	"context"
	"errors"
	"fmt"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"gt-blog/backend/internal/ddd/content_management_context/L3_infrastructure/persistence/mysql/dao"

	"gorm.io/gorm"
)

type MySQLCategoryRepository struct {
	db *gorm.DB
}

func NewMySQLCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &MySQLCategoryRepository{db: db}
}

func (r *MySQLCategoryRepository) Add(ctx context.Context, category *model.Category) (int, error) {
	categoryDAO := r.toDAO(category)
	if err := r.db.WithContext(ctx).Create(categoryDAO).Error; err != nil {
		return 0, err
	}
	category.ID = categoryDAO.ID
	category.Version = categoryDAO.Version
	return category.ID, nil
}

func (r *MySQLCategoryRepository) Update(ctx context.Context, category *model.Category) error {
	if category.ID <= 0 {
		return errors.New("分类ID必须大于0")
	}

	categoryDAO := r.toDAO(category)
	result := r.db.WithContext(ctx).
		Model(&dao.CategoryDAO{}).
		Where("id = ?", category.ID).
		Updates(categoryDAO)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("分类不存在或版本冲突(ID:%d)", category.ID)
	}
	category.Version++
	return nil
}

func (r *MySQLCategoryRepository) FindByID(ctx context.Context, id int) (*model.Category, error) {
	var categoryDAO dao.CategoryDAO
	if err := r.db.WithContext(ctx).First(&categoryDAO, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return r.toDomain(ctx, &categoryDAO), nil
}

func (r *MySQLCategoryRepository) DeleteByID(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("分类ID必须大于0")
	}

	result := r.db.WithContext(ctx).Delete(&dao.CategoryDAO{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("分类不存在(ID:%d)", id)
	}
	return nil
}

func (r *MySQLCategoryRepository) FindByPage(ctx context.Context) ([]*model.Category, int, error) {
	var total int64
	if err := r.db.WithContext(ctx).Model(&dao.CategoryDAO{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var categoryDAOs []*dao.CategoryDAO
	err := r.db.WithContext(ctx).
		Order("create_time DESC").
		Find(&categoryDAOs).Error

	if err != nil {
		return nil, 0, err
	}

	categories := make([]*model.Category, 0, len(categoryDAOs))
	for _, dao := range categoryDAOs {
		categories = append(categories, r.toDomain(ctx, dao))
	}
	return categories, int(total), nil
}

// 领域模型 → 数据模型
func (r *MySQLCategoryRepository) toDAO(category *model.Category) *dao.CategoryDAO {
	return &dao.CategoryDAO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreateTime:  category.CreateTime,
		UpdateTime:  category.UpdateTime,
		Version:     category.Version,
	}
}

// 数据模型 → 领域模型
func (r *MySQLCategoryRepository) toDomain(ctx context.Context, categoryDAO *dao.CategoryDAO) *model.Category {
	// 查询该分类下的文章数量
	var articleCount int64
	r.db.WithContext(ctx).Model(&dao.ArticleDAO{}).
		Where("category_id = ?", categoryDAO.ID).
		Count(&articleCount)

	return &model.Category{
		ID:           categoryDAO.ID,
		Name:         categoryDAO.Name,
		Description:  categoryDAO.Description,
		ArticleCount: int(articleCount),
		CreateTime:   categoryDAO.CreateTime,
		UpdateTime:   categoryDAO.UpdateTime,
		Version:      categoryDAO.Version,
	}
}
