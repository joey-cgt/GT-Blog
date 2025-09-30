package service

import (
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
