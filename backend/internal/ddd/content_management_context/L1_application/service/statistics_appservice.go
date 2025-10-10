package service

import (
	"context"
	"fmt"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"time"
)

type BlogStatAppService struct {
	articleRepo    repository.ArticleRepository
	categoryRepo   repository.CategoryRepository
	tagRepo        repository.TagRepository
	columnRepo     repository.ColumnRepository
	statisticsRepo repository.StatisticsRepository
}

func NewBlogStatAppsService(articleRepo repository.ArticleRepository, categoryRepo repository.CategoryRepository, tagRepo repository.TagRepository, columnRepo repository.ColumnRepository, statisticsRepo repository.StatisticsRepository) *BlogStatAppService {
	return &BlogStatAppService{
		articleRepo:    articleRepo,
		categoryRepo:   categoryRepo,
		tagRepo:        tagRepo,
		columnRepo:     columnRepo,
		statisticsRepo: statisticsRepo,
	}
}

func (s *BlogStatAppService) GetBlogOverviewStatistics(ctx context.Context) (*model.BlogOverviewStats, error) {
	totalArticles, err := s.articleRepo.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	totalLikes, err := s.articleRepo.CountTotalLikes(ctx)
	if err != nil {
		return nil, err
	}
	totalViews, err := s.articleRepo.CountTotalViews(ctx)
	if err != nil {
		return nil, err
	}
	totalColumns, err := s.columnRepo.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	totalCategories, err := s.categoryRepo.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	totalTags, err := s.tagRepo.CountTotal(ctx)
	if err != nil {
		return nil, err
	}
	return &model.BlogOverviewStats{
		TotalArticles:   totalArticles,
		TotalLikes:      totalLikes,
		TotalViews:      totalViews,
		TotalColumns:    totalColumns,
		TotalCategories: totalCategories,
		TotalTags:       totalTags,
	}, nil
}

func (s *BlogStatAppService) GetViewCountTrend(ctx context.Context, days int) ([]*model.ViewCountTrendItem, error) {
	// 调用统计仓库获取趋势数据
	return s.statisticsRepo.GetViewCountTrend(ctx, days)
}

// AggregateDailyData 聚合每日统计数据
func (s *BlogStatAppService) AggregateDailyData(ctx context.Context) error {
	// 获取昨天的日期
	yesterday := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 截断到天

	// 获取总点赞数
	totalLikes, err := s.articleRepo.CountTotalLikes(ctx)
	if err != nil {
		return err
	}

	// 获取总浏览量
	totalViews, err := s.articleRepo.CountTotalViews(ctx)
	if err != nil {
		return err
	}

	// 统计当天新增文章数
	// 直接在应用服务层实现统计逻辑，避免修改仓库接口
	// 查询昨天创建的文章数量
	newArticles, err := s.statisticsRepo.CountNewArticlesByDate(ctx, yesterday)
	if err != nil {
		return fmt.Errorf("统计新增文章数失败: %w", err)
	}

	// 创建博客总统计数据
	blogStats := &model.BlogDailyStats{
		Date:        yesterday,
		TotalViews:  totalViews,
		TotalLikes:  totalLikes,
		NewArticles: newArticles,
	}

	// 保存博客总统计数据
	return s.statisticsRepo.SaveBlogDailyStats(ctx, blogStats)
}

// RecordArticleView 记录文章浏览量
func (s *BlogStatAppService) RecordArticleView(ctx context.Context, articleID int, ip, userAgent, referer string) error {
	// 1. 记录浏览量增量
	if err := s.statisticsRepo.RecordArticleViewIncrement(ctx, articleID, ip, userAgent, referer); err != nil {
		return fmt.Errorf("记录浏览量增量失败: %w", err)
	}

	// 2. 更新文章浏览量
	article, err := s.articleRepo.FindByID(ctx, articleID)
	if err != nil {
		return fmt.Errorf("查找文章失败: %w", err)
	}

	article.IncrementViewCount()
	if err := s.articleRepo.Update(ctx, article); err != nil {
		return fmt.Errorf("更新文章浏览量失败: %w", err)
	}

	return nil
}

// CreateDailyStatsTask 创建每日统计任务
func (s *BlogStatAppService) CreateDailyStatsTask() {
	// 在实际应用中，这里应该设置定时任务，每天凌晨执行AggregateDailyData方法
	// 例如使用github.com/robfig/cron库

	// 示例代码：每天凌晨2点执行聚合任务
	// c := cron.New()
	// c.AddFunc("0 2 * * *", func() {
	//   ctx := context.Background()
	//   if err := s.AggregateDailyData(ctx); err != nil {
	//     // 记录错误日志
	//   }
	// })
	// c.Start()
}
