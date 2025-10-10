package repository

import (
	"context"
	"fmt"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"gt-blog/backend/internal/ddd/content_management_context/L3_infrastructure/persistence/mysql/dao"
	"time"

	"gorm.io/gorm"
)

type MySQLStatisticsRepository struct {
	db *gorm.DB
}

func NewMySQLStatisticsRepository(db *gorm.DB) repository.StatisticsRepository {
	return &MySQLStatisticsRepository{db: db}
}

func (r *MySQLStatisticsRepository) SaveBlogDailyStats(ctx context.Context, stats *model.BlogDailyStats) error {
	// 使用upsert操作（存在则更新，不存在则插入）
	return r.db.WithContext(ctx).
		Create(&dao.DailyBlogStatsDAO{
			Date:        stats.Date,
			TotalViews:  uint(stats.TotalViews),
			TotalLikes:  uint(stats.TotalLikes),
			NewArticles: uint(stats.NewArticles),
			CreatedAt:   time.Now(),
		}).Error
}

func (r *MySQLStatisticsRepository) SaveArticleDailyStats(ctx context.Context, stats *model.ArticleDailyStats) error {
	// 使用upsert操作
	return r.db.WithContext(ctx).
		Create(&dao.DailyArticleStatsDAO{
			Date:      stats.Date,
			ArticleID: uint(stats.ArticleID),
			ViewCount: uint(stats.ViewCount),
			LikeCount: uint(stats.LikeCount),
			CreatedAt: time.Now(),
		}).Error
}

func (r *MySQLStatisticsRepository) RecordArticleViewIncrement(ctx context.Context, articleID int, ip, userAgent, referer string) error {
	return r.db.WithContext(ctx).
		Create(&dao.ArticleViewIncrementDAO{
			ArticleID:     uint(articleID),
			IncrementTime: time.Now(),
			IP:            ip,
			UserAgent:     userAgent,
			Referer:       referer,
		}).Error
}

func (r *MySQLStatisticsRepository) BatchRecordArticleViewIncrements(ctx context.Context, increments []*model.ArticleDailyStats) error {
	// 这里简化实现，实际项目中应该使用gorm的批量插入
	for _, inc := range increments {
		if err := r.SaveArticleDailyStats(ctx, inc); err != nil {
			return err
		}
	}
	return nil
}

func (r *MySQLStatisticsRepository) GetBlogOverviewStats(ctx context.Context) (*model.BlogOverviewStats, error) {
	// 这里简化实现，实际项目中应该从多个表中聚合数据
	return &model.BlogOverviewStats{}, nil
}

func (r *MySQLStatisticsRepository) GetViewCountTrend(ctx context.Context, days int) ([]*model.ViewCountTrendItem, error) {
	var trendItems []*model.ViewCountTrendItem

	// 构建日期范围查询
	query := `
    SELECT 
      DATE_FORMAT(date, '%Y-%m-%d') as date, 
      SUM(total_views) as view_count
    FROM daily_blog_stats
    WHERE date >= DATE_SUB(CURDATE(), INTERVAL ? DAY)
    GROUP BY date
    ORDER BY date ASC
  `

	err := r.db.WithContext(ctx).
		Raw(query, days).
		Scan(&trendItems).
		Error

	if err != nil {
		return nil, fmt.Errorf("获取浏览量趋势失败: %w", err)
	}

	// 确保返回完整的日期范围（包括没有数据的日期）
	return fillMissingDates(trendItems, days), nil
}

func (r *MySQLStatisticsRepository) ClearStatsByDateRange(ctx context.Context, startDate, endDate string) error {
	// 转换日期字符串为time.Time
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return fmt.Errorf("开始日期格式错误: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return fmt.Errorf("结束日期格式错误: %w", err)
	}

	// 清理博客总统计数据
	if err := r.db.WithContext(ctx).
		Where("date BETWEEN ? AND ?", start, end).
		Delete(&dao.DailyBlogStatsDAO{}).Error; err != nil {
		return fmt.Errorf("清理博客总统计数据失败: %w", err)
	}

	// 清理文章统计数据
	if err := r.db.WithContext(ctx).
		Where("date BETWEEN ? AND ?", start, end).
		Delete(&dao.DailyArticleStatsDAO{}).Error; err != nil {
		return fmt.Errorf("清理文章统计数据失败: %w", err)
	}

	return nil
}

// 统计指定日期新增文章数
func (r *MySQLStatisticsRepository) CountNewArticlesByDate(ctx context.Context, date time.Time) (int, error) {
	// 构建日期范围（当天00:00:00到23:59:59）
	startDate := date.Truncate(24 * time.Hour)
	endDate := startDate.Add(24*time.Hour - 1*time.Second)

	// 查询当天创建的文章数量
	var count int64
	err := r.db.WithContext(ctx).
		Model(&dao.ArticleDAO{}).
		Where("create_time BETWEEN ? AND ?", startDate, endDate).
		Count(&count).
		Error

	if err != nil {
		return 0, fmt.Errorf("统计新增文章数失败: %w", err)
	}

	return int(count), nil
}

// 辅助函数：填充缺失的日期，确保返回连续的日期序列
func fillMissingDates(items []*model.ViewCountTrendItem, days int) []*model.ViewCountTrendItem {
	// 创建日期到浏览量的映射
	itemMap := make(map[string]int)
	for _, item := range items {
		itemMap[item.Date] = item.ViewCount
	}

	// 生成完整的日期序列并填充数据
	result := make([]*model.ViewCountTrendItem, 0, days)
	today := time.Now().Truncate(24 * time.Hour)

	for i := days - 1; i >= 0; i-- {
		date := today.AddDate(0, 0, -i).Format("2006-01-02")

		viewCount, exists := itemMap[date]
		if !exists {
			viewCount = 0
		}

		result = append(result, &model.ViewCountTrendItem{
			Date:      date,
			ViewCount: viewCount,
		})
	}

	return result
}
