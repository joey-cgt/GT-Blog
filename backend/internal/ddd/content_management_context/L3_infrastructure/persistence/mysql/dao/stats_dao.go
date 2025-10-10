package dao

import "time"

// DailyBlogStatsDAO 博客每日统计总表DAO
type DailyBlogStatsDAO struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Date        time.Time `gorm:"column:date;not null;uniqueIndex:idx_date"`
	TotalViews  uint      `gorm:"column:total_views;not null;default:0"`
	TotalLikes  uint      `gorm:"column:total_likes;not null;default:0"`
	NewArticles uint      `gorm:"column:new_articles;not null;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;index:idx_created_at"`
}

func (d *DailyBlogStatsDAO) TableName() string {
	return "daily_blog_stats"
}

// DailyArticleStatsDAO 文章每日统计明细表DAO
type DailyArticleStatsDAO struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Date      time.Time `gorm:"column:date;not null;uniqueIndex:idx_date_article_id"`
	ArticleID uint      `gorm:"column:article_id;not null;uniqueIndex:idx_date_article_id"`
	ViewCount uint      `gorm:"column:view_count;not null;default:0"`
	LikeCount uint      `gorm:"column:like_count;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (d *DailyArticleStatsDAO) TableName() string {
	return "daily_article_stats"
}

// ArticleViewIncrementDAO 文章浏览量增量记录DAO
type ArticleViewIncrementDAO struct {
	ID            uint      `gorm:"column:id;primaryKey;autoIncrement"`
	ArticleID     uint      `gorm:"column:article_id;not null;index:idx_article_id"`
	IncrementTime time.Time `gorm:"column:increment_time;not null;index:idx_increment_time"`
	IP            string    `gorm:"column:ip;type:varchar(45)"`
	UserAgent     string    `gorm:"column:user_agent;type:text"`
	Referer       string    `gorm:"column:referer;type:text"`
}

func (d *ArticleViewIncrementDAO) TableName() string {
	return "article_view_increments"
}
