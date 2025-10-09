package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
)

type ArticleRepository interface {
	Add(ctx context.Context, article *model.Article) (int, error)
	Update(ctx context.Context, article *model.Article) error
	FindByID(ctx context.Context, id int) (*model.Article, error)
	DeleteByID(ctx context.Context, id int) error
	FindByPage(ctx context.Context, status int, offset, limit int, keyword string) ([]*model.Article, int, error)
	FindTopArticles(ctx context.Context, limit int) ([]*model.Article, error)
	FindLatestArticles(ctx context.Context, limit int) ([]*model.Article, error)
	FindPopularArticles(ctx context.Context, limit int) ([]*model.Article, error)
	FindByCategoryID(ctx context.Context, categoryID int, offset, limit int) ([]*model.Article, int, error)
	FindByColumnID(ctx context.Context, categoryID int, offset, limit int) ([]*model.Article, int, error)
	FindByTagID(ctx context.Context, categoryID int, offset, limit int) ([]*model.Article, int, error)
}
