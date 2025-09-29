package service

import (
	"context"
	"errors"
	"fmt"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
)

type ArticleDomainService struct {
	articleRepo    repository.ArticleRepository
	articleTagRepo repository.ArticleTagRepository
	tagRepo        repository.TagRepository
	columnRepo     repository.ColumnRepository
	categoryRepo   repository.CategoryRepository
}

// 构造函数：初始化所有依赖的仓储
func NewArticleDomainService(
	articleRepo repository.ArticleRepository,
	articleTagRepo repository.ArticleTagRepository,
	tagRepo repository.TagRepository,
	columnRepo repository.ColumnRepository,
	categoryRepo repository.CategoryRepository,
) *ArticleDomainService {
	return &ArticleDomainService{
		articleRepo:    articleRepo,
		articleTagRepo: articleTagRepo,
		tagRepo:        tagRepo,
		columnRepo:     columnRepo,
		categoryRepo:   categoryRepo,
	}
}

func (s *ArticleDomainService) PublishArticle(ctx context.Context, article *model.Article) (int, error) {
	// 1. 执行文章发布（聚合根自身状态变更）
	if err := article.Publish(); err != nil {
		return 0, err
	}
	// 2. 保存文章（获取文章ID）
	articleID, err := s.articleRepo.Add(ctx, article)
	if err != nil {
		return 0, err
	}
	// 3. 直接更新文章-标签关联表（核心步骤，与发布强相关）
	if len(article.TagIDs) > 0 {
		err := s.articleTagRepo.SaveArticleTags(ctx, articleID, article.TagIDs)
		if err != nil {
			return articleID, err
		}
	}
	// 4. 同步更新标签的文章计数（与关联表更新同属标签相关逻辑）
	if err := s.updateTagArticleCount(ctx, article.TagIDs); err != nil {
		return articleID, err
	}
	// 5. 同步更新专栏的文章计数
	if err := s.updateColumnArticleCount(ctx, article.ColumnID); err != nil {
		return articleID, err
	}
	// 6. 同步更新分类的文章计数
	if err := s.updateCategoryArticleCount(ctx, article.CategoryID); err != nil {
		return articleID, err
	}

	return articleID, nil

}

func (s *ArticleDomainService) UpdateAggregationCounts(ctx context.Context, article *model.Article, isFirstPublish bool) error {
	// 第一次发布文章，直接更新对应的计数
	if isFirstPublish {
		if err := s.updateColumnArticleCount(ctx, article.ColumnID); err != nil {
			return fmt.Errorf("增加新专栏计数失败: %w", err)
		}
		if err := s.updateCategoryArticleCount(ctx, article.CategoryID); err != nil {
			return fmt.Errorf("增加新分类计数失败: %w", err)
		}
		if err := s.updateTagArticleCount(ctx, article.TagIDs); err != nil {
			return fmt.Errorf("增加新标签计数失败: %w", err)
		}
	} else { // 非第一次发布文章，需要先减去旧的计数，再增加新的计数
		originalArticle, err := s.articleRepo.FindByID(ctx, article.ID)
		if err != nil {
			return fmt.Errorf("查询原始文章失败: %w", err)
		}
		if originalArticle == nil {
			return errors.New("文章不存在")
		}

		// 处理标签变更（如果有）
		if !areTagIDsEqual(originalArticle.TagIDs, article.TagIDs) {
			// 先减少旧标签的文章计数
			if err := s.decrementTagArticleCount(ctx, originalArticle.TagIDs); err != nil {
				return fmt.Errorf("减少旧标签计数失败: %w", err)
			}

			// 更新文章-标签关联表
			if len(article.TagIDs) > 0 {
				if err := s.articleTagRepo.SaveArticleTags(ctx, article.ID, article.TagIDs); err != nil {
					return fmt.Errorf("更新标签关联失败: %w", err)
				}
			} else {
				// 如果新标签为空，删除关联
				if err := s.articleTagRepo.DeleteArticleTagsByArticleID(ctx, article.ID); err != nil {
					return fmt.Errorf("删除标签关联失败: %w", err)
				}
			}

			// 增加新标签的文章计数
			if err := s.updateTagArticleCount(ctx, article.TagIDs); err != nil {
				return fmt.Errorf("增加新标签计数失败: %w", err)
			}
		}

		// 处理专栏变更（如果有）
		if originalArticle.ColumnID != article.ColumnID {
			// 减少旧专栏的文章计数
			if originalArticle.ColumnID > 0 {
				if err := s.decrementColumnArticleCount(ctx, originalArticle.ColumnID); err != nil {
					return fmt.Errorf("减少旧专栏计数失败: %w", err)
				}
			}

			// 增加新专栏的文章计数
			if err := s.updateColumnArticleCount(ctx, article.ColumnID); err != nil {
				return fmt.Errorf("增加新专栏计数失败: %w", err)
			}
		}

		// 处理分类变更（如果有）
		if originalArticle.CategoryID != article.CategoryID {
			// 减少旧分类的文章计数
			if originalArticle.CategoryID > 0 {
				if err := s.decrementCategoryArticleCount(ctx, originalArticle.CategoryID); err != nil {
					return fmt.Errorf("减少旧分类计数失败: %w", err)
				}
			}

			// 增加新分类的文章计数
			if err := s.updateCategoryArticleCount(ctx, article.CategoryID); err != nil {
				return fmt.Errorf("增加新分类计数失败: %w", err)
			}
		}

	}

	return nil
}

func (s *ArticleDomainService) updateCategoryArticleCount(ctx context.Context, categoryID int) error {
	if categoryID <= 0 {
		return nil // 未指定分类，无需更新
	}

	// 查询分类
	category, err := s.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return fmt.Errorf("查询分类失败 (ID: %d): %w", categoryID, err)
	}

	if category == nil {
		return fmt.Errorf("分类不存在 (ID: %d)", categoryID)
	}

	// 调用分类聚合根的方法更新计数
	if err := category.IncrementArticleCount(); err != nil {
		return fmt.Errorf("更新分类计数失败 (ID: %d): %w", categoryID, err)
	}

	// 保存更新后的分类
	if err := s.categoryRepo.Update(ctx, category); err != nil {
		return fmt.Errorf("保存分类失败 (ID: %d): %w", categoryID, err)
	}

	return nil
}

func (s *ArticleDomainService) updateColumnArticleCount(ctx context.Context, columnID int) error {
	if columnID <= 0 {
		return nil // 未指定专栏，无需更新
	}

	// 查询专栏
	column, err := s.columnRepo.FindByID(ctx, columnID)
	if err != nil {
		return fmt.Errorf("查询专栏失败 (ID: %d): %w", columnID, err)
	}

	if column == nil {
		return fmt.Errorf("专栏不存在 (ID: %d)", columnID)
	}

	// 调用专栏聚合根的方法更新计数
	if err := column.IncrementArticleCount(); err != nil {
		return fmt.Errorf("更新专栏计数失败 (ID: %d): %w", columnID, err)
	}

	// 保存更新后的专栏
	if err := s.columnRepo.Update(ctx, column); err != nil {
		return fmt.Errorf("保存专栏失败 (ID: %d): %w", columnID, err)
	}

	return nil
}

func (s *ArticleDomainService) updateTagArticleCount(ctx context.Context, tagIDs []int) error {
	if len(tagIDs) == 0 {
		return nil // 没有标签，无需更新
	}
	// 查询并更新每个标签的文章计数
	for _, tagID := range tagIDs {
		tag, err := s.tagRepo.FindByID(ctx, tagID)
		if err != nil {
			return fmt.Errorf("查询标签失败 (ID: %d): %w", tagID, err)
		}

		if tag == nil {
			return fmt.Errorf("标签不存在 (ID: %d)", tagID)
		}

		// 调用标签聚合根的方法更新计数（保持聚合根自治）
		if err := tag.IncrementArticleCount(); err != nil {
			return fmt.Errorf("更新标签计数失败 (ID: %d): %w", tagID, err)
		}

		// 保存更新后的标签
		if err := s.tagRepo.Update(ctx, tag); err != nil {
			return fmt.Errorf("保存标签失败 (ID: %d): %w", tagID, err)
		}
	}
	return nil
}

// 辅助方法：减少标签的文章计数
func (s *ArticleDomainService) decrementTagArticleCount(ctx context.Context, tagIDs []int) error {
	for _, tagID := range tagIDs {
		tag, err := s.tagRepo.FindByID(ctx, tagID)
		if err != nil {
			return fmt.Errorf("查询标签失败 (ID: %d): %w", tagID, err)
		}

		if tag == nil {
			continue // 标签已被删除，跳过
		}

		if err := tag.DecrementArticleCount(); err != nil {
			return fmt.Errorf("减少标签计数失败 (ID: %d): %w", tagID, err)
		}

		if err := s.tagRepo.Update(ctx, tag); err != nil {
			return fmt.Errorf("保存标签失败 (ID: %d): %w", tagID, err)
		}
	}

	return nil
}

// 辅助方法：减少专栏的文章计数
func (s *ArticleDomainService) decrementColumnArticleCount(ctx context.Context, columnID int) error {
	column, err := s.columnRepo.FindByID(ctx, columnID)
	if err != nil {
		return fmt.Errorf("查询专栏失败 (ID: %d): %w", columnID, err)
	}

	if column == nil {
		return nil // 专栏已被删除，跳过
	}

	if err := column.DecrementArticleCount(); err != nil {
		return fmt.Errorf("减少专栏计数失败 (ID: %d): %w", columnID, err)
	}

	if err := s.columnRepo.Update(ctx, column); err != nil {
		return fmt.Errorf("保存专栏失败 (ID: %d): %w", columnID, err)
	}

	return nil
}

// 辅助方法：减少分类的文章计数
func (s *ArticleDomainService) decrementCategoryArticleCount(ctx context.Context, categoryID int) error {
	category, err := s.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return fmt.Errorf("查询分类失败 (ID: %d): %w", categoryID, err)
	}

	if category == nil {
		return nil // 分类已被删除，跳过
	}

	if err := category.DecrementArticleCount(); err != nil {
		return fmt.Errorf("减少分类计数失败 (ID: %d): %w", categoryID, err)
	}

	if err := s.categoryRepo.Update(ctx, category); err != nil {
		return fmt.Errorf("保存分类失败 (ID: %d): %w", categoryID, err)
	}

	return nil
}

// 辅助函数：比较两个标签ID切片是否相等
func areTagIDsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// 使用map比较元素
	aMap := make(map[int]bool)
	for _, id := range a {
		aMap[id] = true
	}

	for _, id := range b {
		if !aMap[id] {
			return false
		}
		delete(aMap, id)
	}

	return len(aMap) == 0
}
