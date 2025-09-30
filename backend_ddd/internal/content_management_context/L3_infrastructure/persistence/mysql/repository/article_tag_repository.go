package repository

import (
	"context"
	"errors"
	"fmt"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
	"gt-blog/backend_ddd/internal/content_management_context/L3_infrastructure/persistence/mysql/dao"

	"gorm.io/gorm"
)

type MySQLArticleTagRepository struct {
	db *gorm.DB
}

func NewMySQLArticleTagRepository(db *gorm.DB) repository.ArticleTagRepository {
	return &MySQLArticleTagRepository{db: db}
}

// DeleteArticleTagsByArticleID implements repository.ArticleTagRepository.
func (m *MySQLArticleTagRepository) DeleteArticleTagsByArticleID(ctx context.Context, articleID int) error {
	if articleID <= 0 {
		return errors.New("文章ID必须大于0")
	}

	result := m.db.WithContext(ctx).
		Where("article_id = ?", articleID).
		Delete(&dao.ArticleTagDAO{})

	if result.Error != nil {
		return fmt.Errorf("删除文章标签关联失败: %w", result.Error)
	}
	return nil
}

// DeleteArticleTagsByTagID implements repository.ArticleTagRepository.
func (m *MySQLArticleTagRepository) DeleteArticleTagsByTagID(ctx context.Context, tagID int) error {
	if tagID <= 0 {
		return errors.New("标签ID必须大于0")
	}

	result := m.db.WithContext(ctx).
		Where("tag_id = ?", tagID).
		Delete(&dao.ArticleTagDAO{})

	if result.Error != nil {
		return fmt.Errorf("删除标签文章关联失败: %w", result.Error)
	}
	return nil
}

// SaveArticleTags implements repository.ArticleTagRepository.
func (m *MySQLArticleTagRepository) SaveArticleTags(ctx context.Context, articleID int, tagIDs []int) error {
	if articleID <= 0 {
		return errors.New("文章ID必须大于0")
	}
	if len(tagIDs) == 0 {
		return nil
	}

	// 构建批量插入数据
	var articleTags []*dao.ArticleTagDAO
	for _, tagID := range tagIDs {
		if tagID <= 0 {
			return errors.New("标签ID必须大于0")
		}
		articleTags = append(articleTags, &dao.ArticleTagDAO{
			ArticleID: articleID,
			TagID:     tagID,
		})
	}

	// 批量插入
	if err := m.db.WithContext(ctx).Create(&articleTags).Error; err != nil {
		return fmt.Errorf("保存文章标签关联失败: %w", err)
	}
	return nil
}

func (m *MySQLArticleTagRepository) UpdateArticleTags(ctx context.Context, articleID int, oldTagIDs, newTagIDs []int) error {
	if articleID <= 0 {
		return errors.New("文章ID必须大于0")
	}
	if len(oldTagIDs) == 0 && len(newTagIDs) == 0 {
		return nil
	}

	// 计算需要删除的标签ID（存在于旧标签列表但不存在于新标签列表）
	var tagsToDelete []int
	oldTagMap := make(map[int]bool)
	for _, id := range oldTagIDs {
		oldTagMap[id] = true
	}
	for _, id := range oldTagIDs {
		found := false
		for _, newID := range newTagIDs {
			if id == newID {
				found = true
				break
			}
		}
		if !found {
			tagsToDelete = append(tagsToDelete, id)
		}
	}

	// 计算需要新增的标签ID（存在于新标签列表但不存在于旧标签列表）
	var tagsToAdd []int
	for _, newID := range newTagIDs {
		if !oldTagMap[newID] {
			tagsToAdd = append(tagsToAdd, newID)
		}
	}

	// 执行删除操作
	if len(tagsToDelete) > 0 {
		if err := m.DeleteArticleTagsByArticleTagIDs(ctx, articleID, tagsToDelete); err != nil {
			return fmt.Errorf("删除文章标签关联失败: %w", err)
		}
	}

	// 执行新增操作
	if len(tagsToAdd) > 0 {
		if err := m.SaveArticleTags(ctx, articleID, tagsToAdd); err != nil {
			return fmt.Errorf("保存文章标签关联失败: %w", err)
		}
	}
	return nil
}

func (m *MySQLArticleTagRepository) DeleteArticleTagsByArticleTagIDs(ctx context.Context, articleID int, tagsToDelete []int) error {
	if articleID <= 0 {
		return errors.New("文章ID必须大于0")
	}
	if len(tagsToDelete) == 0 {
		return nil
	}

	result := m.db.WithContext(ctx).
		Where("article_id = ? AND tag_id IN ?", articleID, tagsToDelete).
		Delete(&dao.ArticleTagDAO{})
	if result.Error != nil {
		return fmt.Errorf("删除文章标签关联失败: %w", result.Error)
	}
	return nil
}
