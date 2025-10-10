package repository

import (
	"context"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"time"
)

type StatisticsRepository interface {
	// 保存博客每日总统计数据（如果已存在则更新）
	SaveBlogDailyStats(ctx context.Context, stats *model.BlogDailyStats) error

	// 保存文章每日统计数据（如果已存在则更新）
	SaveArticleDailyStats(ctx context.Context, stats *model.ArticleDailyStats) error

	// 记录文章浏览量增量
	RecordArticleViewIncrement(ctx context.Context, articleID int, ip, userAgent, referer string) error

	// 批量记录文章浏览量增量（用于性能优化）
	BatchRecordArticleViewIncrements(ctx context.Context, increments []*model.ArticleDailyStats) error

	// 获取博客总览统计数据
	GetBlogOverviewStats(ctx context.Context) (*model.BlogOverviewStats, error)

	// 获取浏览量趋势数据（按天聚合）
	GetViewCountTrend(ctx context.Context, days int) ([]*model.ViewCountTrendItem, error)

	// 清空指定日期范围的统计数据（用于重新计算）
	ClearStatsByDateRange(ctx context.Context, startDate, endDate string) error

	// 统计指定日期新增文章数
	CountNewArticlesByDate(ctx context.Context, date time.Time) (int, error)
}
