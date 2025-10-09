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

type MySQLColumnRepository struct {
	db *gorm.DB
}

func NewMySQLColumnRepository(db *gorm.DB) repository.ColumnRepository {
	return &MySQLColumnRepository{db: db}
}

// Add implements repository.ColumnRepository.
func (m *MySQLColumnRepository) Add(ctx context.Context, column *model.Column) (int, error) {
	columnDAO := m.toDAO(column)
	if err := m.db.WithContext(ctx).Create(columnDAO).Error; err != nil {
		return 0, err
	}
	column.ID = columnDAO.ID
	column.Version = columnDAO.Version
	return column.ID, nil
}

// FindByID implements repository.ColumnRepository.
func (m *MySQLColumnRepository) FindByID(ctx context.Context, id int) (*model.Column, error) {
	var columnDAO dao.ColumnDAO
	if err := m.db.WithContext(ctx).First(&columnDAO, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return m.toDomain(ctx, &columnDAO), nil
}

// Update implements repository.ColumnRepository.
func (m *MySQLColumnRepository) Update(ctx context.Context, column *model.Column) error {
	if column.ID <= 0 {
		return errors.New("专栏ID必须大于0")
	}

	columnDAO := m.toDAO(column)
	result := m.db.WithContext(ctx).
		Model(&dao.ColumnDAO{}).
		Where("id = ?", column.ID).
		Updates(columnDAO)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("专栏不存在或版本冲突(ID:%d)", column.ID)
	}
	column.Version++
	return nil
}

// DeleteByID implements repository.ColumnRepository.
func (m *MySQLColumnRepository) DeleteByID(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("专栏ID必须大于0")
	}

	result := m.db.WithContext(ctx).Delete(&dao.ColumnDAO{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("专栏不存在(ID:%d)", id)
	}
	return nil
}

// FindByPage implements repository.ColumnRepository.
func (m *MySQLColumnRepository) FindByPage(ctx context.Context, page, pageSize int) ([]*model.Column, int, error) {
	var total int64
	if err := m.db.WithContext(ctx).Model(&dao.ColumnDAO{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var columnDAOs []*dao.ColumnDAO
	err := m.db.WithContext(ctx).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("create_time DESC").
		Find(&columnDAOs).Error

	if err != nil {
		return nil, 0, err
	}

	columns := make([]*model.Column, 0, len(columnDAOs))
	for _, columnDAO := range columnDAOs {
		columns = append(columns, m.toDomain(ctx, columnDAO))
	}
	return columns, int(total), nil
}

// 领域模型 → 数据模型
func (m *MySQLColumnRepository) toDAO(column *model.Column) *dao.ColumnDAO {
	return &dao.ColumnDAO{
		ID:          column.ID,
		Name:        column.Name,
		Description: column.Description,
		CreateTime:  column.CreateTime,
		UpdateTime:  column.UpdateTime,
		Version:     column.Version,
	}
}

// 数据模型 → 领域模型
func (m *MySQLColumnRepository) toDomain(ctx context.Context, columnDAO *dao.ColumnDAO) *model.Column {
	// 查询该专栏下的文章数量
	var articleCount int64
	m.db.WithContext(ctx).Model(&dao.ArticleDAO{}).
		Where("column_id = ?", columnDAO.ID).
		Count(&articleCount)

	return &model.Column{
		ID:           columnDAO.ID,
		Name:         columnDAO.Name,
		Description:  columnDAO.Description,
		ArticleCount: articleCount,
		CreateTime:   columnDAO.CreateTime,
		UpdateTime:   columnDAO.UpdateTime,
		Version:      columnDAO.Version,
	}
}
