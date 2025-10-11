package service

import (
	"context"
	"time"

	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
)

type CommentAppService struct {
	commentRepo repository.CommentRepository
}

func NewCommentAppService(commentRepo repository.CommentRepository) *CommentAppService {
	return &CommentAppService{
		commentRepo: commentRepo,
	}
}

// SubmitComment 提交评论
func (s *CommentAppService) SubmitComment(ctx context.Context, req request.CommentSubmitReq) error {
	// 创建评论模型
	comment := &model.Comment{
		Username:  req.Nickname,
		Email:     req.Email,
		Content:   req.Content,
		ArticleID: uint(req.ArticleID),
		ParentID:  nil, // 默认是顶级评论
		CreatedAt: time.Now(),
	}

	// 调用仓库保存评论
	return s.commentRepo.CreateComment(ctx, comment)
}

// GetArticleComments 获取文章的评论列表
func (s *CommentAppService) GetArticleComments(ctx context.Context, articleID uint) (result.CommentsResult, error) {
	comments, err := s.commentRepo.GetCommentsByArticleID(ctx, articleID)
	if err != nil {
		return result.CommentsResult{}, err
	}
	// 将comments转成CommentsResult
	var commentResult result.CommentsResult
	// 只调用一次buildCommentTree，避免重复计算
	commentTree := buildCommentTree(comments)
	commentResult.TotalCount = countCommentItems(commentTree)
	commentResult.Comments = commentTree
	return commentResult, nil
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

// buildCommentTree 构建评论树结构
func buildCommentTree(comments []model.Comment) []result.CommentItemResult {
	// 创建评论ID到评论的映射
	commentMap := make(map[uint]*result.CommentItemResult)
	var rootComments []result.CommentItemResult

	// 先将所有评论转换为结果模型并存储在映射中
	for _, comment := range comments {
		commentItem := &result.CommentItemResult{
			ID:        comment.ID,
			Username:  comment.Username,
			Email:     comment.Email,
			Content:   comment.Content,
			ArticleID: comment.ArticleID,
			ParentID:  comment.ParentID,
			Children:  make([]result.CommentItemResult, 0), // 初始化Children字段
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		commentMap[comment.ID] = commentItem
	}

	// 构建评论树结构
	for _, comment := range comments {
		commentItem := commentMap[comment.ID]

		// 如果是顶级评论（没有父评论），直接添加到结果中
		if comment.ParentID == nil {
			rootComments = append(rootComments, *commentItem)
		} else {
			// 如果是子评论，添加到父评论的Children列表中
			parentID := *comment.ParentID
			if parentComment, exists := commentMap[parentID]; exists {
				parentComment.Children = append(parentComment.Children, *commentItem)
			}
		}
	}

	return rootComments
}
