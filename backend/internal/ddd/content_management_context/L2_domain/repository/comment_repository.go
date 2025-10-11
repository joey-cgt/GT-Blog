package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
)

type CommentRepository interface {
	// CreateComment 创建评论
	CreateComment(ctx context.Context, comment *model.Comment) error
	// GetCommentsByArticleID 获取文章的评论列表
	GetCommentsByArticleID(ctx context.Context, articleID uint) ([]model.Comment, error)
	// DeleteComment 删除评论
	DeleteComment(ctx context.Context, commentID uint) error
}
