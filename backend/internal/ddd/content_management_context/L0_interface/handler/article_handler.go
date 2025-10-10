package handler

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/converter"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/query"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/service"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleAppService *service.ArticleAppService
}

func NewArticleHandler(articleAppService *service.ArticleAppService) *ArticleHandler {
	return &ArticleHandler{
		articleAppService: articleAppService,
	}
}

// 创建并发布文章
func (h *ArticleHandler) CreateAndPublishArticle(c *gin.Context) {
	var req request.CreateAndPublishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertCreateAndPublishRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	articleId, err := h.articleAppService.CreateAndPublishArticle(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "创建文章失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"articleId": articleId})
}

// 更新已发布文章
func (h *ArticleHandler) UpdatePublishedArticle(c *gin.Context) {
	var req request.UpdatePublishedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdatePublishedRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.articleAppService.UpdatePublishedArticle(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新文章失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "更新文章成功"})
}

// 创建并保存为草稿
func (h *ArticleHandler) CreateDraft(c *gin.Context) {
	var req request.CreateDraftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertCreateDraftRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	// 调用服务层保存草稿
	draftId, err := h.articleAppService.CreateDraft(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "保存草稿失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"draftId": draftId})
}

// 更新草稿
func (h *ArticleHandler) UpdateDraft(c *gin.Context) {
	var req request.UpdateDraftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdateDraftRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.articleAppService.UpdateDraft(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新草稿失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "更新草稿成功"})
}

// 发布草稿
func (h *ArticleHandler) PublishDraft(c *gin.Context) {
	var req request.PublishDraftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertPublishDraftRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if _, err := h.articleAppService.PublishDraft(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "发布草稿失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "发布草稿成功"})
}

// 删除单篇文章
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	var req request.DeleteArticleRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertDeleteArticleRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.articleAppService.DeleteArticle(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "删除文章失败: " + err.Error()})
	}

	c.JSON(200, gin.H{"message": "删除文章成功"})
}

// 设置文章置顶状态
func (h *ArticleHandler) UpdateArticleTopStatus(c *gin.Context) {
	var req request.UpdateArticleTopStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdateArticleTopStatusRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.articleAppService.UpdateArticleTopStatus(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新文章置顶状态失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "更新文章置顶状态成功"})
}

// 获取单篇文章详情
func (h *ArticleHandler) GetArticleDetail(c *gin.Context) {
	var req request.GetArticleDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数格式错误" + err.Error()})
		return
	}

	qry, err := converter.ConvertGetArticleDetailRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "转换参数格式错误" + err.Error()})
		return
	}

	articleResult, err := h.articleAppService.GetArticleDetail(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取文章详情失败: " + err.Error()})
		return
	}

	resp := converter.ConvertArticleDetailResultToResponse(articleResult)

	c.JSON(200, gin.H{"data": resp})
}

// 查看全部文章
func (h *ArticleHandler) GetAllArticleList(c *gin.Context) {
	var req request.GetAllArticleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	qry, err := converter.ConvertGetAllArticleListRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	articleListResult, err := h.articleAppService.GetAllArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取文章列表失败: " + err.Error()})
		return
	}

	resp := converter.ConvertAllArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})

}

// 查看置顶文章
func (h *ArticleHandler) GetTopArticleList(c *gin.Context) {
	qry := query.GetHomePageArticleListQuery{
		Type:      "top",
		SortBy:    "update_time",
		SortOrder: "desc",
	}
	articleListResult, err := h.articleAppService.GetHomePageArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取置顶文章列表失败: " + err.Error()})
		return
	}
	resp := converter.ConvertHomePageArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})
}

// 查看热门文章
func (h *ArticleHandler) GetPopularArticleList(c *gin.Context) {
	qry := query.GetHomePageArticleListQuery{
		Type:      "hot",
		SortBy:    "update_time",
		SortOrder: "desc",
	}
	articleListResult, err := h.articleAppService.GetHomePageArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取热门文章列表失败: " + err.Error()})
		return
	}
	resp := converter.ConvertHomePageArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})
}

// 查看最新文章
func (h *ArticleHandler) GetLatestArticleList(c *gin.Context) {
	qry := query.GetHomePageArticleListQuery{
		Type:      "new",
		SortBy:    "update_time",
		SortOrder: "desc",
	}
	articleListResult, err := h.articleAppService.GetHomePageArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取最新文章列表失败: " + err.Error()})
		return
	}
	resp := converter.ConvertHomePageArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})
}

func (h *ArticleHandler) GetMostViewedArticleList(c *gin.Context) {
	qry := query.GetSortedArticleListQuery{
		Limit:     5,
		SortBy:    "view_count",
		SortOrder: "desc",
	}
	articleListResult, err := h.articleAppService.GetSortedArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取最多查看文章列表失败: " + err.Error()})
		return
	}
	resp := converter.ConvertSortedArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})
}

func (h *ArticleHandler) GetMostLikedArticleList(c *gin.Context) {
	qry := query.GetSortedArticleListQuery{
		Limit:     5,
		SortBy:    "like_count",
		SortOrder: "desc",
	}
	articleListResult, err := h.articleAppService.GetSortedArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取最多点赞文章列表失败: " + err.Error()})
		return
	}
	resp := converter.ConvertSortedArticleListResultToResponse(articleListResult)

	c.JSON(200, gin.H{"data": resp})
}

func (h *ArticleHandler) GetAggregatedArticleList(c *gin.Context) {
	var req request.GetAggregatedArticleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	qry, err := converter.ConvertGetAggregatedArticleListRequestToQuery(req)

	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}
	articleResult, err := h.articleAppService.GetAggregatedArticleList(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取分类文章失败: " + err.Error()})
		return
	}
	resp := converter.ConvertAggregatedArticleListResultToResponse(articleResult)

	c.JSON(200, gin.H{"data": resp})
}

func (h *ArticleHandler) GetArticleListByKeyword(c *gin.Context) {

}

func (h *ArticleHandler) IncrementLike(c *gin.Context) {
	var req request.IncrementLikeRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数格式错误" + err.Error()})
		return
	}
	if err := h.articleAppService.IncrementLike(c, req.ID); err != nil {
		c.JSON(500, gin.H{"error": "增加点赞数失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "增加点赞数成功"})
}

func (h *ArticleHandler) DecrementLike(c *gin.Context) {
	var req request.DecrementLikeRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "请求参数格式错误" + err.Error()})
		return
	}
	if err := h.articleAppService.DecrementLike(c, req.ID); err != nil {
		c.JSON(500, gin.H{"error": "减少点赞数失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "减少点赞数成功"})
}
