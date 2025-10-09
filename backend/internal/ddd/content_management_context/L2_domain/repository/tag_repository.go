package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
)

type TagRepository interface {
	Add(ctx context.Context, tag *model.Tag) (int, error)
	Update(ctx context.Context, tag *model.Tag) error
	FindByID(ctx context.Context, id int) (*model.Tag, error)
	FindByIDs(ctx context.Context, ids []int) ([]*model.Tag, error)
	IncrementArticleCount(ctx context.Context, id int) error
	DeleteByID(ctx context.Context, id int) error
	FindByPage(ctx context.Context, page, pageSize int) ([]*model.Tag, int, error)
}
