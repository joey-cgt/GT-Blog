package repository

import (
	"context"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
)

type ColumnRepository interface {
	Add(ctx context.Context, column *model.Column) (int, error)
	Update(ctx context.Context, column *model.Column) error
	FindByID(ctx context.Context, id int) (*model.Column, error)
	DeleteByID(ctx context.Context, id int) error
	FindByPage(ctx context.Context, page, pageSize int) ([]*model.Column, int, error)
}
