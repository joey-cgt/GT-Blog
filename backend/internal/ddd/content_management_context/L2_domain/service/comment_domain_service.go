package service

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
)

type CommentDomainService struct {
	commentRepo repository.CommentRepository
}

func NewCommentDomainService(commentRepo repository.CommentRepository) *CommentDomainService {
	return &CommentDomainService{
		commentRepo: commentRepo,
	}
}

func (s *CommentDomainService) GetCommentsByArticleID(ctx context.Context, articleID uint) ([]model.Comment, error) {
	comments, err := s.commentRepo.GetCommentsByArticleID(ctx, articleID)
	if err != nil {
		return nil, err
	}
	commentMap := make(map[uint]*model.Comment)
	for _, comment := range comments {
		commentMap[comment.ID] = comment
	}
	// 构建评论树
	var commentTree []*model.Comment
	for _, comment := range comments {
		if comment.ParentID == nil {
			commentTree = append(commentTree, comment)
		} else {
			parent, exists := commentMap[*comment.ParentID]
			if exists {
				parent.Children = append(parent.Children, comment)
			}
		}
	}
	// 转换为Comment类型
	var commentResult []model.Comment
	for _, comment := range commentTree {
		commentResult = append(commentResult, *comment)
	}

	return commentResult, nil
}
