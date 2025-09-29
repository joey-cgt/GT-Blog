package service

import (
	"context"
	"fmt"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/command"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/query"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/result"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/service"
	"math"
)

type ArticleAppService struct {
	articleDomainService *service.ArticleDomainService
	articleRepo          repository.ArticleRepository
	tagRepo              repository.TagRepository
	columnRepo           repository.ColumnRepository
	categoryRepo         repository.CategoryRepository
}

func NewArticleAppService(
	articleDomainService *service.ArticleDomainService,
	articleRepo repository.ArticleRepository,
	tagRepo repository.TagRepository,
	columnRepo repository.ColumnRepository,
	categoryRepo repository.CategoryRepository,
) *ArticleAppService {
	return &ArticleAppService{
		articleDomainService: articleDomainService,
		articleRepo:          articleRepo,
		tagRepo:              tagRepo,
		columnRepo:           columnRepo,
		categoryRepo:         categoryRepo,
	}
}

func (s *ArticleAppService) CreateAndPublishArticle(ctx context.Context, cmd command.CreateAndPublishCommand) (int, error) {
	// 新建文章
	newArticle, err := model.NewArticle(
		cmd.Title,
		cmd.Abstract,
		cmd.Content,
		cmd.CoverUrl,
		cmd.ColumnID,
		cmd.CategoryID,
		cmd.TagIDs,
	)
	if err != nil {
		return 0, err
	}

	// 执行文章发布（聚合根自身状态变更）
	err = newArticle.Publish()
	if err != nil {
		return 0, err
	}

	// 先持久化
	id, err := s.articleRepo.Add(ctx, newArticle)
	if err != nil {
		return 0, err
	}

	// 再更新相关计数
	article, err := s.articleRepo.FindByID(ctx, id)
	if err != nil {
		return 0, err
	}
	if err := s.articleDomainService.UpdateAggregationCounts(ctx, article, true); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ArticleAppService) UpdatePublishedArticle(ctx context.Context, cmd command.UpdatePublishedCommand) error {
	// 获取文章
	article, err := s.articleRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if article.Status != 1 {
		return fmt.Errorf("文章不存在或未发布")
	}
	// 先更新文章相关计数
	if err := s.articleDomainService.UpdateAggregationCounts(ctx, article, false); err != nil {
		return err
	}

	// 再更新文章并持久化
	if err := article.Update(
		cmd.Title,
		cmd.Abstract,
		cmd.Content,
		cmd.ColumnID,
		cmd.CategoryID,
		cmd.TagIDs,
		cmd.CoverUrl,
	); err != nil {
		return err
	}
	// 检查文章发布状态，防止更新时内容不合法
	if err := article.Publish(); err != nil {
		return err
	}
	// 最后再持久化
	return s.articleRepo.Update(ctx, article)

}

func (s *ArticleAppService) CreateDraft(ctx context.Context, cmd command.CreateDraftCommand) (int, error) {
	newDraft, err := model.NewArticle(
		cmd.Title,
		cmd.Abstract,
		cmd.Content,
		cmd.CoverUrl,
		cmd.ColumnID,
		cmd.CategoryID,
		cmd.TagIDs,
	)
	if err != nil {
		return 0, err
	}

	return s.articleRepo.Add(ctx, newDraft)
}

func (s *ArticleAppService) UpdateDraft(ctx context.Context, cmd command.UpdateDraftCommand) error {
	draft, err := s.articleRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if draft.Status != 0 {
		return fmt.Errorf("文章不是草稿状态")
	}

	if err := draft.Update(
		cmd.Title,
		cmd.Abstract,
		cmd.Content,
		cmd.ColumnID,
		cmd.CategoryID,
		cmd.TagIDs,
		cmd.CoverUrl,
	); err != nil {
		return err
	}

	return s.articleRepo.Update(ctx, draft)
}

func (s *ArticleAppService) PublishDraft(ctx context.Context, cmd command.PublishDraftCommand) (int, error) {
	draft, err := s.articleRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return 0, err
	}

	if err := draft.Update(
		cmd.Title,
		cmd.Abstract,
		cmd.Content,
		cmd.ColumnID,
		cmd.CategoryID,
		cmd.TagIDs,
		cmd.CoverUrl,
	); err != nil {
		return 0, err
	}

	if err := draft.Publish(); err != nil {
		return 0, err
	}

	if err := s.articleDomainService.UpdateAggregationCounts(ctx, draft, true); err != nil {
		return 0, err
	}

	return draft.ID, nil
}

func (s *ArticleAppService) DeleteArticle(ctx context.Context, cmd command.DeleteArticleCommand) error {
	err := s.articleRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleAppService) UpdateArticleTopStatus(ctx context.Context, cmd command.UpdateArticleTopStatusCommand) error {
	article, err := s.articleRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	article.UpdateTopStatus(cmd.IsTop)

	return s.articleRepo.Update(ctx, article)
}

func (s *ArticleAppService) GetArticleDetail(ctx context.Context, qry query.GetArticleDetailQuery) (*result.ArticleDetailResult, error) {
	article, err := s.articleRepo.FindByID(ctx, qry.ID)
	if err != nil {
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}

	if article == nil {
		return nil, fmt.Errorf("文章不存在 (ID: %d)", qry.ID)
	}

	categoryResult := s.buildCategoryResult(ctx, article.CategoryID)
	columnResult := s.buildColumnResult(ctx, article.ColumnID)
	tagResults := s.buildTagResults(ctx, article.TagIDs)

	articleResult := &result.ArticleDetailResult{
		Title:       article.Title,
		Content:     article.Content,
		Abstract:    article.Abstract,
		CoverUrl:    article.CoverUrl,
		Status:      article.Status,
		IsTop:       article.IsTop,
		Category:    categoryResult,
		Column:      columnResult,
		Tags:        tagResults,
		LikeCount:   article.LikeCount,
		ViewCount:   article.ViewCount,
		CreateTime:  article.CreateTime.Format("2006-01-02 15:04:05"),
		PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  article.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	return articleResult, nil
}

func (s *ArticleAppService) GetAllArticleList(ctx context.Context, qry query.GetAllArticleListQuery) (*result.AllArticleListResult, error) {
	if qry.PageSize <= 0 {
		qry.PageSize = 10
	}
	if qry.Page <= 0 {
		qry.Page = 1
	}
	offset := (qry.Page - 1) * qry.PageSize
	articles, total, err := s.articleRepo.FindByPage(ctx, qry.Status, offset, qry.PageSize, qry.Keyword)
	if err != nil {
		return nil, err
	}

	items := make([]*result.ArticleListItemResult, 0, len(articles))
	for _, article := range articles {

		categoryResult := s.buildCategoryResult(ctx, article.CategoryID)
		columnResult := s.buildColumnResult(ctx, article.ColumnID)
		tagResults := s.buildTagResults(ctx, article.TagIDs)
		items = append(items, &result.ArticleListItemResult{
			ID:          article.ID,
			Title:       article.Title,
			Abstract:    article.Abstract,
			Category:    categoryResult,
			Column:      columnResult,
			Tags:        tagResults,
			ViewCount:   article.ViewCount,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			IsTop:       article.IsTop,
		})
	}
	// 计算总页数
	totalPages := int(math.Ceil(float64(total) / float64(qry.PageSize)))

	return &result.AllArticleListResult{
		Items:      items,
		Total:      total,
		Page:       qry.Page,
		PageSize:   qry.PageSize,
		TotalPages: totalPages,
	}, nil

}

func (s *ArticleAppService) GetHomePageArticleList(ctx context.Context, qry query.GetHomePageArticleListQuery) (*result.HomePageArticleListResult, error) {
	var articleList []*model.Article
	var err error
	switch qry.Type {
	case "top":
		articleList, err = s.articleRepo.FindTopArticles(ctx, 5)
	case "new":
		articleList, err = s.articleRepo.FindLatestArticles(ctx, 5)
	case "hot":
		articleList, err = s.articleRepo.FindPopularArticles(ctx, 5)
	default:
		return nil, fmt.Errorf("未知请求")
	}

	if err != nil {
		return nil, err
	}

	// 如果没有置顶文章，返回空列表
	if len(articleList) == 0 {
		return &result.HomePageArticleListResult{
			Items:     []*result.ArticleListItemResult{},
			Type:      qry.Type,
			SortBy:    qry.SortBy,
			SortOrder: qry.SortOrder,
		}, nil
	}

	// 构建响应结果
	items := make([]*result.ArticleListItemResult, 0, len(articleList))
	for _, article := range articleList {
		categoryResult := s.buildCategoryResult(ctx, article.CategoryID)
		columnResult := s.buildColumnResult(ctx, article.ColumnID)
		tagResults := s.buildTagResults(ctx, article.TagIDs)

		items = append(items, &result.ArticleListItemResult{
			ID:          article.ID,
			Title:       article.Title,
			Abstract:    article.Abstract,
			Category:    categoryResult,
			Column:      columnResult,
			Tags:        tagResults,
			ViewCount:   article.ViewCount,
			LikeCount:   article.LikeCount,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			IsTop:       article.IsTop,
		})
	}

	return &result.HomePageArticleListResult{
		Items:     items,
		Type:      qry.Type,
		SortBy:    qry.SortBy,
		SortOrder: qry.SortOrder,
	}, nil
}

func (s *ArticleAppService) GetAggregatedArticleList(ctx context.Context, qry query.GetAggregatedArticleListQuery) (*result.AggregatedArticleListResult, error) {
	if qry.PageSize <= 0 {
		qry.PageSize = 10
	}
	if qry.Page <= 0 {
		qry.Page = 1
	}
	offset := (qry.Page - 1) * qry.PageSize

	var articleList []*model.Article
	var err error
	var total int
	switch qry.Type {
	case "categpry":
		articleList, total, err = s.articleRepo.FindByCategoryID(ctx, qry.ID, offset, qry.PageSize)
	case "column":
		articleList, total, err = s.articleRepo.FindByColumnID(ctx, qry.ID, offset, qry.PageSize)
	case "tag":
		articleList, total, err = s.articleRepo.FindByTagID(ctx, qry.ID, offset, qry.PageSize)
	default:
		return nil, fmt.Errorf("未知请求")
	}

	if err != nil {
		return nil, err
	}

	if len(articleList) == 0 {
		return &result.AggregatedArticleListResult{
			Items:    []*result.ArticleListItemResult{},
			Type:     qry.Type,
			Total:    0,
			Page:     0,
			PageSize: 0,
		}, nil
	}

	// 构建响应结果
	items := make([]*result.ArticleListItemResult, 0, len(articleList))
	for _, article := range articleList {
		categoryResult := s.buildCategoryResult(ctx, article.CategoryID)
		columnResult := s.buildColumnResult(ctx, article.ColumnID)
		tagResults := s.buildTagResults(ctx, article.TagIDs)

		items = append(items, &result.ArticleListItemResult{
			ID:          article.ID,
			Title:       article.Title,
			Abstract:    article.Abstract,
			Category:    categoryResult,
			Column:      columnResult,
			Tags:        tagResults,
			ViewCount:   article.ViewCount,
			LikeCount:   article.LikeCount,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			IsTop:       article.IsTop,
		})
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(total) / float64(qry.PageSize)))

	return &result.AggregatedArticleListResult{
		Items:      items,
		Type:       qry.Type,
		Total:      total,
		Page:       qry.Page,
		PageSize:   qry.PageSize,
		TotalPages: totalPages,
	}, nil
}

// ==========================================================================================================
// 辅助方法：构建分类结果
func (s *ArticleAppService) buildCategoryResult(ctx context.Context, categoryID int) *result.CategoryResult {
	if categoryID <= 0 {
		return nil
	}
	category, err := s.categoryRepo.FindByID(ctx, categoryID)
	if err != nil || category == nil {
		return nil
	}
	return &result.CategoryResult{
		ID:   category.ID,
		Name: category.Name,
	}
}

// 辅助方法：构建专栏结果
func (s *ArticleAppService) buildColumnResult(ctx context.Context, columnID int) *result.ColumnResult {
	if columnID <= 0 {
		return nil
	}
	column, err := s.columnRepo.FindByID(ctx, columnID)
	if err != nil || column == nil {
		return nil
	}
	return &result.ColumnResult{
		ID:   column.ID,
		Name: column.Name,
	}
}

// 辅助方法：构建标签结果（批量查询优化）
func (s *ArticleAppService) buildTagResults(ctx context.Context, tagIDs []int) []*result.TagResult {
	if len(tagIDs) == 0 {
		return nil
	}
	tags, err := s.tagRepo.FindByIDs(ctx, tagIDs)
	if err != nil || len(tags) == 0 {
		return nil
	}
	tagResults := make([]*result.TagResult, 0, len(tags))
	for _, tag := range tags {
		tagResults = append(tagResults, &result.TagResult{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	return tagResults
}
