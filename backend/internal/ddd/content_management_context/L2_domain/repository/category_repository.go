package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
)

type CategoryRepository interface {
	Add(ctx context.Context, category *model.Category) (int, error)
	Update(ctx context.Context, category *model.Category) error
	FindByID(ctx context.Context, id int) (*model.Category, error)
	DeleteByID(ctx context.Context, id int) error
	FindByPage(ctx context.Context) ([]*model.Category, int, error)
	CountTotal(ctx context.Context) (int, error)
}
