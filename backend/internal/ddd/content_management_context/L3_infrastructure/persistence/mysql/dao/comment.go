package dao

import "time"

type CommentDAO struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null;index"`     // 评论者用户名
	Email    string `gorm:"not null;index"`     // 评论者邮箱
	Content  string `gorm:"type:text;not null"` // 评论内容
	// 关联字段 (外键)
	ArticleID uint `gorm:"not null;index"` // 关联的文章ID
	// 自引用，实现回复功能
	ParentID  *uint        `gorm:"index"`               // 父评论ID。为NULL表示是顶级评论，否则是回复。
	Children  []CommentDAO `gorm:"foreignKey:ParentID"` // 子评论列表 (用于查询)
	CreatedAt time.Time
}
