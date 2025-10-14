package service

import (
	"context"
	"fmt"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/command"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/query"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/service"
	"math"
)

type ArticleAppService struct {
	articleDomainService *service.ArticleDomainService
	articleRepo          repository.ArticleRepository
	tagRepo              repository.TagRepository
	columnRepo           repository.ColumnRepository
	categoryRepo         repository.CategoryRepository
	articleTagRepo       repository.ArticleTagRepository
	commentRepo          repository.CommentRepository
}

func NewArticleAppService(
	articleDomainService *service.ArticleDomainService,
	articleRepo repository.ArticleRepository,
	tagRepo repository.TagRepository,
	columnRepo repository.ColumnRepository,
	categoryRepo repository.CategoryRepository,
	articleTagRepo repository.ArticleTagRepository,
	commentRepo repository.CommentRepository,
) *ArticleAppService {
	return &ArticleAppService{
		articleDomainService: articleDomainService,
		articleRepo:          articleRepo,
		tagRepo:              tagRepo,
		columnRepo:           columnRepo,
		categoryRepo:         categoryRepo,
		articleTagRepo:       articleTagRepo,
		commentRepo:          commentRepo,
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

	// 更新 article_tags 表
	if len(newArticle.TagIDs) > 0 {
		err = s.articleTagRepo.SaveArticleTags(ctx, id, newArticle.TagIDs)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (s *ArticleAppService) UpdatePublishedArticle(ctx context.Context, cmd command.UpdatePublishedCommand) error {
	// 获取文章
	article, err := s.articleRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	oldTagIDs := article.TagIDs

	if article.Status != 1 {
		return fmt.Errorf("文章不存在或未发布")
	}

	// 更新文章并持久化
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
	// 持久化
	if err := s.articleRepo.Update(ctx, article); err != nil {
		return err
	}

	newTagIDs := article.TagIDs

	if err := s.articleTagRepo.UpdateArticleTags(ctx, article.ID, oldTagIDs, newTagIDs); err != nil {
		return err
	}
	return nil
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

	if len(draft.TagIDs) > 0 {
		err := s.articleTagRepo.SaveArticleTags(ctx, draft.ID, draft.TagIDs)
		if err != nil {
			return draft.ID, err
		}
	}

	// 最后再持久化
	if err := s.articleRepo.Update(ctx, draft); err != nil {
		return draft.ID, err
	}

	return draft.ID, nil
}

func (s *ArticleAppService) DeleteArticle(ctx context.Context, cmd command.DeleteArticleCommand) error {
	err := s.articleRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := s.articleTagRepo.DeleteArticleTagsByArticleID(ctx, cmd.ID); err != nil {
		return err
	}

	// 级联删除相关评论
	comments, err := s.commentRepo.GetCommentsByArticleID(ctx, uint(cmd.ID))
	if err != nil {
		return fmt.Errorf("获取文章评论失败: %w", err)
	}
	
	// 遍历删除所有评论
	for _, comment := range comments {
		if err := s.commentRepo.DeleteComment(ctx, comment.ID); err != nil {
			return fmt.Errorf("删除评论失败(ID: %d): %w", comment.ID, err)
		}
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

	// 构建响应结果
	items := s.buildArticleListItemResults(ctx, articles)
	// 单独设置创建时间（部分列表需要）
	for i, item := range items {
		item.CreateTime = articles[i].CreateTime.Format("2006-01-02 15:04:05")
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
	items := s.buildArticleListItemResults(ctx, articleList)

	return &result.HomePageArticleListResult{
		Items:     items,
		Type:      qry.Type,
		SortBy:    qry.SortBy,
		SortOrder: qry.SortOrder,
	}, nil
}

func (s *ArticleAppService) GetSortedArticleList(ctx context.Context, qry query.GetSortedArticleListQuery) (*result.SortedArticleListResult, error) {
	if qry.Limit <= 0 {
		qry.Limit = 5
	}
	articles, err := s.articleRepo.FindSortedArticles(ctx, qry.SortBy, qry.SortOrder, qry.Limit)
	if err != nil {
		return nil, err
	}

	// 构建响应结果
	items := s.buildArticleListItemResults(ctx, articles)

	return &result.SortedArticleListResult{
		Items:     items,
		Limit:     qry.Limit,
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
	case "category":
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
	items := s.buildArticleListItemResults(ctx, articleList)

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

func (s *ArticleAppService) IncrementLike(ctx context.Context, id int) error {
	return s.articleRepo.IncrementLikeById(ctx, id)
}

func (s *ArticleAppService) DecrementLike(ctx context.Context, id int) error {
	return s.articleRepo.DecrementLikeById(ctx, id)
}

// GetRecommendedArticles 获取推荐文章列表
// 推荐规则：
// 1. 优先查询同一专栏的文章
// 2. 如果文章有标签，查询相同标签的文章
// 3. 如果以上两个方式查找的文章还不够limit篇，再查找同一个分类下的文章
// 4. 最终的文章数量不超过limit篇
func (s *ArticleAppService) GetRecommendedArticles(ctx context.Context, articleID int, limit int) (*result.RecommendedArticleListResult, error) {
	// 参数校验
	if articleID <= 0 {
		return nil, fmt.Errorf("文章ID必须大于0")
	}
	if limit <= 0 {
		limit = 5 // 默认返回5篇
	}

	// 获取当前文章详情
	sourceArticle, err := s.articleRepo.FindByID(ctx, articleID)
	if err != nil || sourceArticle == nil {
		return nil, fmt.Errorf("获取文章详情失败: %w", err)
	}

	// 存储推荐文章的ID集合，用于去重
	recommendedIDs := make(map[int]bool)
	recommendedIDs[articleID] = true // 排除当前文章

	// 存储推荐文章列表
	var recommendedArticles []*model.Article

	// 1. 优先查询同一专栏的文章
	if sourceArticle.ColumnID > 0 {
		columnArticles, _, err := s.articleRepo.FindByColumnID(ctx, sourceArticle.ColumnID, 0, limit)
		if err == nil {
			for _, article := range columnArticles {
				if !recommendedIDs[article.ID] && len(recommendedArticles) < limit {
					recommendedIDs[article.ID] = true
					recommendedArticles = append(recommendedArticles, article)
				}
			}
		}
	}

	// 2. 如果还需要更多文章且当前文章有标签，查询相同标签的文章
	if len(recommendedArticles) < limit && len(sourceArticle.TagIDs) > 0 {
		// 遍历所有标签，获取相关文章
		for _, tagID := range sourceArticle.TagIDs {
			if len(recommendedArticles) >= limit {
				break
			}
			tagArticles, _, err := s.articleRepo.FindByTagID(ctx, tagID, 0, limit-len(recommendedArticles))
			if err == nil {
				for _, article := range tagArticles {
					if !recommendedIDs[article.ID] && len(recommendedArticles) < limit {
						recommendedIDs[article.ID] = true
						recommendedArticles = append(recommendedArticles, article)
					}
				}
			}
		}
	}

	// 3. 如果还需要更多文章，查询同一分类下的文章
	if len(recommendedArticles) < limit && sourceArticle.CategoryID > 0 {
		categoryArticles, _, err := s.articleRepo.FindByCategoryID(ctx, sourceArticle.CategoryID, 0, limit-len(recommendedArticles))
		if err == nil {
			for _, article := range categoryArticles {
				if !recommendedIDs[article.ID] && len(recommendedArticles) < limit {
					recommendedIDs[article.ID] = true
					recommendedArticles = append(recommendedArticles, article)
				}
			}
		}
	}

	// 构建响应结果
	items := s.buildArticleListItemResults(ctx, recommendedArticles)

	return &result.RecommendedArticleListResult{
		Items: items,
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

// 辅助方法：构建文章列表项结果
func (s *ArticleAppService) buildArticleListItemResults(ctx context.Context, articles []*model.Article) []*result.ArticleListItemResult {
	items := make([]*result.ArticleListItemResult, 0, len(articles))
	for _, article := range articles {
		categoryResult := s.buildCategoryResult(ctx, article.CategoryID)
		columnResult := s.buildColumnResult(ctx, article.ColumnID)
		tagResults := s.buildTagResults(ctx, article.TagIDs)

		items = append(items, &result.ArticleListItemResult{
			ID:          article.ID,
			Title:       article.Title,
			Abstract:    article.Abstract,
			CoverUrl:    article.CoverUrl,
			Category:    categoryResult,
			Column:      columnResult,
			Tags:        tagResults,
			ViewCount:   article.ViewCount,
			LikeCount:   article.LikeCount,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			IsTop:       article.IsTop,
		})
	}
	return items
}
