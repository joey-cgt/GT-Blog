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
func (m *MySQLCommentRepository) GetCommentsByArticleID(ctx context.Context, articleID uint) ([]*model.Comment, error) {
	// 查询所有属于该文章的评论
	var allComments []dao.CommentDAO
	if err := m.db.
		WithContext(ctx).
		Where("article_id = ?", articleID).
		Order("created_at ASC").
		Find(&allComments).Error; err != nil {
		return nil, err
	}

	// 将评论转换为领域模型
	var allCommentDomain []*model.Comment
	for _, commentDAO := range allComments {
		comment := m.toDomain(&commentDAO)
		allCommentDomain = append(allCommentDomain, comment)
	}
	return allCommentDomain, nil
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

// toDomain converts a CommentDAO to a Comment domain model.
func (m *MySQLCommentRepository) toDomain(commentDAO *dao.CommentDAO) *model.Comment {
	return &model.Comment{
		ID:        commentDAO.ID,
		Nickname:  commentDAO.Nickname,
		Email:     commentDAO.Email,
		Content:   commentDAO.Content,
		ArticleID: commentDAO.ArticleID,
		ParentID:  commentDAO.ParentID,
		Children:  make([]*model.Comment, 0), // 初始化为空切片
		CreatedAt: commentDAO.CreatedAt,
	}
}

// toDAO converts a Comment domain model to a CommentDAO.
func (m *MySQLCommentRepository) toDAO(comment *model.Comment) *dao.CommentDAO {
	return &dao.CommentDAO{
		ID:        comment.ID,
		Nickname:  comment.Nickname,
		Email:     comment.Email,
		Content:   comment.Content,
		ArticleID: comment.ArticleID,
		ParentID:  comment.ParentID,
		CreatedAt: comment.CreatedAt,
	}
}
