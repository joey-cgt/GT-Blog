package model

import "time"

// BlogOverviewStats 博客总览统计数据模型
type BlogOverviewStats struct {
	TotalArticles   int
	TotalLikes      int
	TotalViews      int
	TotalColumns    int
	TotalCategories int
	TotalTags       int
}

// ViewCountTrendItem 浏览量趋势项
type ViewCountTrendItem struct {
	Date      string
	ViewCount int
}

// ArticleDailyStats 文章每日统计数据模型
type ArticleDailyStats struct {
	ID        int
	Date      time.Time
	ArticleID int
	ViewCount int
	LikeCount int
}

// BlogDailyStats 博客每日总统计数据模型
type BlogDailyStats struct {
	ID          int
	Date        time.Time
	TotalViews  int
	TotalLikes  int
	NewArticles int
}
