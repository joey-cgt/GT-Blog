package router

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/handler"

	"github.com/gin-gonic/gin"
)

// RegisterArticleRoutes 注册文章相关路由
func RegisterArticleRoutes(router *gin.RouterGroup, articleHandler *handler.ArticleHandler) {
	// 文章管理路由组
	articleGroup := router.Group("/articles")
	{
		// 创建并发布文章
		articleGroup.POST("", articleHandler.CreateAndPublishArticle)

		// 更新已发布文章
		articleGroup.PUT("/:id", articleHandler.UpdatePublishedArticle)

		// 保存草稿
		articleGroup.POST("/drafts", articleHandler.CreateDraft)

		// 发布草稿
		articleGroup.POST("/drafts/:id/publish", articleHandler.PublishDraft)

		// 删除文章（已发布或草稿）
		articleGroup.DELETE("/:id", articleHandler.DeleteArticle)

		// 更新文章置顶状态
		articleGroup.PATCH("/:id/top", articleHandler.UpdateArticleTopStatus)

		// 获取文章详情
		articleGroup.GET("/:id", articleHandler.GetArticleDetail)

		// 获取全部文章（已发布、草稿）列表（支持分页、排序、筛选）
		articleGroup.GET("", articleHandler.GetAllArticleList)

		// 获取置顶文章列表
		articleGroup.GET("/top", articleHandler.GetTopArticleList)

		// 获取热门文章列表
		articleGroup.GET("/popular", articleHandler.GetPopularArticleList)

		// 获取最新文章列表
		articleGroup.GET("/latest", articleHandler.GetLatestArticleList)

		// 获取浏览量最高的前五篇文章
		articleGroup.GET("/most-viewed", articleHandler.GetMostViewedArticleList)

		// 获取点赞量最高的前五篇文章
		articleGroup.GET("/most-liked", articleHandler.GetMostLikedArticleList)

		// 获取某专栏/分类/标签下的文章列表
		articleGroup.GET("/aggregated", articleHandler.GetAggregatedArticleList)

		// 关键字搜索文章
		articleGroup.GET("/search", articleHandler.GetArticleListByKeyword)

		articleGroup.POST("/:id/like", articleHandler.IncrementLike)
		articleGroup.DELETE("/:id/like", articleHandler.DecrementLike)
		articleGroup.GET("/recommended", articleHandler.GetRecommendedArticles)
	}
}

func RegisterCategoryRoutes(router *gin.RouterGroup, categoryHandler *handler.CategoryHandler) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("", categoryHandler.GetCategoryList)
		categoryGroup.GET("/:id", categoryHandler.GetCategoryDetail)
		categoryGroup.POST("", categoryHandler.CreateCategory)
		categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}

func RegisterColumnRoutes(router *gin.RouterGroup, columnHandler *handler.ColumnHandler) {
	columnGroup := router.Group("/columns")
	{
		columnGroup.GET("", columnHandler.GetColumnList)
		columnGroup.GET("/:id", columnHandler.GetColumnDetail)
		columnGroup.POST("", columnHandler.CreateColumn)
		columnGroup.PUT("/:id", columnHandler.UpdateColumn)
		columnGroup.DELETE("/:id", columnHandler.DeleteColumn)
	}
}

func RegisterTagRoutes(router *gin.RouterGroup, tagHandler *handler.TagHandler) {
	tagGroup := router.Group("/tags")
	{
		tagGroup.GET("", tagHandler.GetTagList)
		tagGroup.GET("/:id", tagHandler.GetTagDetail)
		tagGroup.POST("", tagHandler.CreateTag)
		tagGroup.PUT("/:id", tagHandler.UpdateTag)
		tagGroup.DELETE("/:id", tagHandler.DeleteTag)
	}
}

func RegisterStatisticsRoutes(router *gin.RouterGroup, statisticsHandler *handler.StatisticsHandler) {
	// 统计数据路由组
	statisticsGroup := router.Group("/statistics")
	{
		// 获取博客总览统计数据
		statisticsGroup.GET("/overview", statisticsHandler.GetBlogOverview)

		// 获取浏览量趋势数据（支持7天、30天、90天）
		statisticsGroup.GET("/view-trend", statisticsHandler.GetViewCountTrend)
	}

	// 记录浏览量的路由
	router.POST("/articles/:id/view", statisticsHandler.RecordView)
}
