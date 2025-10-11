package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L3_infrastructure/persistence/mysql/dao"

	"gorm.io/gorm"
)

type MySQLCommentRepository struct {
	db *gorm.DB
}

func NewMySQLCommentRepository(db *gorm.DB) *MySQLCommentRepository {
	return &MySQLCommentRepository{
		db: db,
	}
}

// CreateComment implements CommentRepository.
func (m *MySQLCommentRepository) CreateComment(ctx context.Context, comment *model.Comment) error {
	commentDAO := m.toDAO(comment)
	return m.db.Create(commentDAO).Error
}

// GetCommentsByArticleID implements CommentRepository.
// 获取文章的评论列表并返回评论树
func (m *MySQLCommentRepository) GetCommentsByArticleID(ctx context.Context, articleID uint) ([]model.Comment, error) {
	// 查询所有属于该文章的评论
	var allComments []dao.CommentDAO
	if err := m.db.
		WithContext(ctx).
		Where("article_id = ?", articleID).
		Order("created_at ASC").
		Find(&allComments).Error; err != nil {
		return nil, err
	}

	// 将评论转换为领域模型并构建评论树
	return m.buildCommentTree(allComments), nil
}

// DeleteComment implements CommentRepository.
func (m *MySQLCommentRepository) DeleteComment(ctx context.Context, commentID uint) error {
	// 先删除子评论
	if err := m.db.
		WithContext(ctx).
		Where("parent_id = ?", commentID).
		Delete(&dao.CommentDAO{}).Error; err != nil {
		return err
	}

	// 再删除父评论
	return m.db.
		WithContext(ctx).
		Where("id = ?", commentID).
		Delete(&dao.CommentDAO{}).Error
}

// DeleteCommentByID 为了向后兼容而保留的方法
func (m *MySQLCommentRepository) DeleteCommentByID(commentID uint) error {
	return m.DeleteComment(context.Background(), commentID)
}

// buildCommentTree 构建评论树结构
func (m *MySQLCommentRepository) buildCommentTree(comments []dao.CommentDAO) []model.Comment {
	// 创建评论ID到评论的映射
	commentMap := make(map[uint]*model.Comment)
	var rootComments []model.Comment

	// 先将所有评论转换为领域模型并存储在映射中
	for _, commentDAO := range comments {
		comment := m.toDomain(&commentDAO)
		commentMap[comment.ID] = comment
	}

	// 构建评论树结构
	for _, commentDAO := range comments {
		comment := commentMap[commentDAO.ID]

		// 如果是顶级评论（没有父评论），直接添加到结果中
		if commentDAO.ParentID == nil {
			rootComments = append(rootComments, *comment)
		} else {
			// 如果是子评论，添加到父评论的Children列表中
			parentID := *commentDAO.ParentID
			if parentComment, exists := commentMap[parentID]; exists {
				parentComment.Children = append(parentComment.Children, *comment)
			}
		}
	}

	return rootComments
}

// toDomain converts a CommentDAO to a Comment domain model.
func (m *MySQLCommentRepository) toDomain(commentDAO *dao.CommentDAO) *model.Comment {
	return &model.Comment{
		ID:        commentDAO.ID,
		Username:  commentDAO.Username,
		Email:     commentDAO.Email,
		Content:   commentDAO.Content,
		ArticleID: commentDAO.ArticleID,
		ParentID:  commentDAO.ParentID,
		Children:  make([]model.Comment, 0), // 初始化为空切片
		CreatedAt: commentDAO.CreatedAt,
	}
}

// toDAO converts a Comment domain model to a CommentDAO.
func (m *MySQLCommentRepository) toDAO(comment *model.Comment) *dao.CommentDAO {
	return &dao.CommentDAO{
		ID:        comment.ID,
		Username:  comment.Username,
		Email:     comment.Email,
		Content:   comment.Content,
		ArticleID: comment.ArticleID,
		ParentID:  comment.ParentID,
		CreatedAt: comment.CreatedAt,
	}
}
