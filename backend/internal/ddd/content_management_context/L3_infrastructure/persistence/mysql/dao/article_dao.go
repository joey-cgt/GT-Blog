package dao

import (
	"time"

	"gorm.io/gorm"
)

// ArticleDAO 仅用于MySQL数据映射，与领域模型字段一一对应，但包含存储细节
type ArticleDAO struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`            // 主键自增（技术特性）
	Title       string    `gorm:"column:title;type:varchar(255);not null"`       // 字段类型约束（技术特性）
	Abstract    string    `gorm:"column:abstract;type:text"`                     // 摘要可为空
	Content     string    `gorm:"column:content;type:text;not null"`             // 内容非空（业务规则→存储约束）
	ColumnID    int       `gorm:"column:column_id;default:0"`                    // 专栏ID默认0（无专栏）
	CategoryID  int       `gorm:"column:category_id;not null"`                   // 分类ID非空（业务规则→存储约束）
	CoverUrl    string    `gorm:"column:cover_url;type:varchar(512)"`            // 封面URL长度限制
	Status      int       `gorm:"column:status;type:tinyint;not null;default:0"` // 状态用tinyint节省空间
	IsTop       bool      `gorm:"column:is_top;type:tinyint;not null;default:0"` // bool在MySQL中存为0/1
	ViewCount   int       `gorm:"column:view_count;not null;default:0"`
	LikeCount   int       `gorm:"column:like_count;not null;default:0"`
	CreateTime  time.Time `gorm:"column:create_time;not null"`
	PublishTime time.Time `gorm:"column:publish_time;default:null"` // 未发布时可为null
	UpdateTime  time.Time `gorm:"column:update_time;not null"`
	Version     int       `gorm:"column:version;not null;default:0"` // 乐观锁版本号
}

// TableName 显式指定数据库表名（技术细节）
func (a *ArticleDAO) TableName() string {
	return "articles"
}

// BeforeSave GORM钩子：自动维护更新时间（技术细节，与业务无关）
func (a *ArticleDAO) BeforeSave(tx *gorm.DB) error {
	a.UpdateTime = time.Now()
	if a.ID == 0 { // 新增时自动设置创建时间
		a.CreateTime = time.Now()
	}
	return nil
}
