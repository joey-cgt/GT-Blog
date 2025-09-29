package repository

import (
	"context"
	"errors"
	"fmt"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
	"gt-blog/backend_ddd/internal/content_management_context/L3_infrastructure/persistence/mysql/dao"

	"gorm.io/gorm"
)

type MySQLTagRepository struct {
	db *gorm.DB
}

func NewMySQLTagRepository(db *gorm.DB) repository.TagRepository {
	return &MySQLTagRepository{db: db}
}

// Add implements repository.TagRepository.
func (m *MySQLTagRepository) Add(ctx context.Context, tag *model.Tag) (int, error) {
	tagDAO := m.toDAO(tag)
	if err := m.db.WithContext(ctx).Create(tagDAO).Error; err != nil {
		return 0, err
	}
	tag.ID = tagDAO.ID
	tag.Version = tagDAO.Version
	return tag.ID, nil
}

// FindByID implements repository.TagRepository.
func (m *MySQLTagRepository) FindByID(ctx context.Context, id int) (*model.Tag, error) {
	var tagDAO dao.TagDAO
	if err := m.db.WithContext(ctx).First(&tagDAO, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return m.toDomain(ctx, &tagDAO), nil
}

func (m *MySQLTagRepository) FindByIDs(ctx context.Context, ids []int) ([]*model.Tag, error) {
	var tagDAOs []*dao.TagDAO
	if err := m.db.WithContext(ctx).Find(&tagDAOs, "id IN ?", ids).Error; err != nil {
		return nil, err
	}

	tags := make([]*model.Tag, 0, len(tagDAOs))
	for _, tagDAO := range tagDAOs {
		tags = append(tags, m.toDomain(ctx, tagDAO))
	}
	return tags, nil
}

func (m *MySQLTagRepository) FindByPage(ctx context.Context, page, pageSize int) ([]*model.Tag, int, error) {
	var tagDAOs []*dao.TagDAO
	err := m.db.WithContext(ctx).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&tagDAOs).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := m.db.WithContext(ctx).Model(&dao.TagDAO{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	tags := make([]*model.Tag, 0, len(tagDAOs))
	for _, tagDAO := range tagDAOs {
		tags = append(tags, m.toDomain(ctx, tagDAO))
	}
	return tags, int(total), nil
}

// IncrementArticleCount implements repository.TagRepository.
func (m *MySQLTagRepository) IncrementArticleCount(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).
		Model(&dao.TagDAO{}).
		Where("id = ?", id).
		Update("article_count", gorm.Expr("article_count + 1"))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("标签不存在(ID:%d)", id)
	}
	return nil
}

// Update implements repository.TagRepository.
func (m *MySQLTagRepository) Update(ctx context.Context, tag *model.Tag) error {
	if tag.ID <= 0 {
		return errors.New("标签ID必须大于0")
	}

	tagDAO := m.toDAO(tag)
	result := m.db.WithContext(ctx).
		Model(&dao.TagDAO{}).
		Where("id = ?", tag.ID).
		Updates(tagDAO)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("标签不存在或版本冲突(ID:%d)", tag.ID)
	}
	tag.Version++
	return nil
}

// DeleteByID implements repository.TagRepository.
func (m *MySQLTagRepository) DeleteByID(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("标签ID必须大于0")
	}

	result := m.db.WithContext(ctx).Delete(&dao.TagDAO{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("标签不存在(ID:%d)", id)
	}
	return nil
}

// 领域模型 → 数据模型
func (m *MySQLTagRepository) toDAO(tag *model.Tag) *dao.TagDAO {
	return &dao.TagDAO{
		ID:          tag.ID,
		Name:        tag.Name,
		Description: tag.Description,
		CreateTime:  tag.CreateTime,
		UpdateTime:  tag.UpdateTime,
		Version:     tag.Version,
	}
}

// 数据模型 → 领域模型
func (m *MySQLTagRepository) toDomain(ctx context.Context, tagDAO *dao.TagDAO) *model.Tag {
	// 查询该标签下的文章数量
	var articleCount int64
	// 先查询关联表获取包含该标签的文章ID数量
	m.db.WithContext(ctx).Model(&dao.ArticleTagDAO{}).
		Where("tag_id = ?", tagDAO.ID).
		Count(&articleCount)

	return &model.Tag{
		ID:           tagDAO.ID,
		Name:         tagDAO.Name,
		Description:  tagDAO.Description,
		ArticleCount: int(articleCount),
		CreateTime:   tagDAO.CreateTime,
		UpdateTime:   tagDAO.UpdateTime,
		Version:      tagDAO.Version,
	}
}
