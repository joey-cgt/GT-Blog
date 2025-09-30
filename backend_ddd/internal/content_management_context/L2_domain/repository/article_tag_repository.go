package repository

import "context"

type ArticleTagRepository interface {
	SaveArticleTags(ctx context.Context, articleID int, tagIDs []int) error
	DeleteArticleTagsByArticleID(ctx context.Context, articleID int) error
	DeleteArticleTagsByTagID(ctx context.Context, tagID int) error
	UpdateArticleTags(ctx context.Context, articleID int, oldTagIDs, newTagIDs []int) error
	DeleteArticleTagsByArticleTagIDs(ctx context.Context, articleID int, tagsToDelete []int) error
}
