package service

import (
	"context"
	"time"

	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/service"
)

type CommentAppService struct {
	commentRepo          repository.CommentRepository
	commentDomainService *service.CommentDomainService
}

func NewCommentAppService(commentRepo repository.CommentRepository, commentDomainService *service.CommentDomainService) *CommentAppService {
	return &CommentAppService{
		commentRepo:          commentRepo,
		commentDomainService: commentDomainService,
	}
}

// SubmitComment 提交评论
func (s *CommentAppService) SubmitComment(ctx context.Context, req request.CommentSubmitReq) error {
	// 创建评论模型
	comment := &model.Comment{
		Nickname:  req.Nickname,
		Email:     req.Email,
		Content:   req.Content,
		ArticleID: uint(req.ArticleID),
		ParentID:  req.ParentID, // 使用请求中的ParentID，如果没有则为nil
		CreatedAt: time.Now(),
	}

	// 调用仓库保存评论
	return s.commentRepo.CreateComment(ctx, comment)
}

// GetArticleComments 获取文章的评论列表
func (s *CommentAppService) GetArticleComments(ctx context.Context, articleID uint) (result.CommentsResult, error) {
	comments, err := s.commentDomainService.GetCommentsByArticleID(ctx, articleID)
	if err != nil {
		return result.CommentsResult{}, err
	}
	var res []result.CommentItemResult
	for _, comment := range comments {
		res = append(res, convertCommentToResult(comment))
	}
	return result.CommentsResult{
		TotalCount: countCommentItems(res),
		Comments:   res,
	}, nil
}

// DeleteCommentByID 删除评论
func (s *CommentAppService) DeleteCommentByID(ctx context.Context, commentID uint) error {
	// 调用仓库删除评论
	return s.commentRepo.DeleteComment(ctx, commentID)
}

// countCommentItems 递归计算评论项数量
func countCommentItems(comments []result.CommentItemResult) int {
	count := 0
	for _, comment := range comments {
		count++
		// 递归计算子评论数量
		if len(comment.Children) > 0 {
			count += countCommentItems(comment.Children)
		}
	}
	return count
}

// convertCommentToResult 递归地将model.Comment及其子评论转换为result.CommentItemResult
func convertCommentToResult(comment model.Comment) result.CommentItemResult {
	resultItem := result.CommentItemResult{
		ID:        comment.ID,
		Nickname:  comment.Nickname,
		Email:     comment.Email,
		Content:   comment.Content,
		ArticleID: comment.ArticleID,
		ParentID:  comment.ParentID,
		Children:  make([]result.CommentItemResult, 0),
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// 递归转换所有子评论
	for _, childComment := range comment.Children {
		resultItem.Children = append(resultItem.Children, convertCommentToResult(*childComment))
	}

	return resultItem
}
