package handler

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/response"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statAppService *service.BlogStatAppService
}

func NewStatisticsHandler(statAppService *service.BlogStatAppService) *StatisticsHandler {
	return &StatisticsHandler{
		statAppService: statAppService,
	}
}

// GetBlogOverview 获取博客总览统计数据
func (h *StatisticsHandler) GetBlogOverview(c *gin.Context) {
	stats, err := h.statAppService.GetBlogOverviewStatistics(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "获取统计数据失败: " + err.Error()})
		return
	}
	resp := response.BlogOverviewStatsResp{
		TotalArticles:   stats.TotalArticles,
		TotalLikes:      stats.TotalLikes,
		TotalViews:      stats.TotalViews,
		TotalCategories: stats.TotalCategories,
		TotalTags:       stats.TotalTags,
	}

	c.JSON(200, gin.H{"data": resp})
}

// GetViewCountTrend 获取指定天数的浏览量趋势
func (h *StatisticsHandler) GetViewCountTrend(c *gin.Context) {
	// 从查询参数获取时间范围，默认为7天
	daysStr := c.DefaultQuery("range", "7d")
	// 将7d 30d 90d 转成int值
	days, err := strconv.Atoi(daysStr[:len(daysStr)-1])
	if err != nil || days <= 0 {
		days = 7
	}

	// 支持的时间范围
	supportedDays := map[int]bool{7: true, 30: true, 90: true}
	if !supportedDays[days] {
		c.JSON(400, gin.H{"error": "不支持的时间范围，支持的时间范围为7天、30天、90天"})
		return
	}

	res, err := h.statAppService.GetViewCountTrend(c, days)
	if err != nil {
		c.JSON(400, gin.H{"error": "获取浏览量趋势失败: " + err.Error()})
		return
	}

	// 将trend转成response.ViewCountTrendItemResp
	trends := make([]response.ViewCountTrendItemResp, 0, len(res))
	for _, item := range res {
		trends = append(trends, response.ViewCountTrendItemResp{
			Date:  item.Date,
			Views: item.ViewCount,
		})
	}

	trendResp := response.ViewCountTrendResp{
		Trend: trends,
		Days:  days,
	}

	c.JSON(200, gin.H{"data": trendResp})
}

// RecordView 记录文章浏览
func (h *StatisticsHandler) RecordView(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "无效的文章ID"})
		return
	}

	// 获取IP、UserAgent和Referer
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	referer := c.Request.Referer()

	// 记录浏览
	if err := h.statAppService.RecordArticleView(c, articleID, ip, userAgent, referer); err != nil {
		// 记录错误但不影响用户体验
		// 实际应用中应该记录日志
	}

	// 无论成功失败都返回成功，不影响用户体验
	c.JSON(http.StatusOK, gin.H{"success": true})
}
