package converter

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/response"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/command"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/query"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"strconv"
)

// ================================== 前端请求转换为后端模型 ==================================

func ConvertCreateAndPublishRequestToCommand(req request.CreateAndPublishRequest) (command.CreateAndPublishCommand, error) {
	cmd := command.CreateAndPublishCommand{
		Title:    req.Title,
		Abstract: req.Abstract,
		Content:  req.Content,
	}

	if req.CoverUrl != nil {
		cmd.CoverUrl = *req.CoverUrl
	}

	// 转换CategoryID从string到int
	categoryID, err := strconv.Atoi(req.CategoryID)
	if err != nil {
		return command.CreateAndPublishCommand{}, err
	}
	cmd.CategoryID = categoryID

	// 转换ColumnID从string到int（可选字段）
	if req.ColumnID != nil {
		columnID, err := strconv.Atoi(*req.ColumnID)
		if err != nil {
			return command.CreateAndPublishCommand{}, err
		}
		cmd.ColumnID = columnID
	}

	// 转换TagIDs从[]string到[]int（可选字段）
	var tagIDs []int
	for _, tagIDStr := range req.TagIDs {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			return command.CreateAndPublishCommand{}, err
		}
		tagIDs = append(tagIDs, tagID)
	}
	cmd.TagIDs = tagIDs

	return cmd, nil
}

func ConvertUpdatePublishedRequestToCommand(req request.UpdatePublishedRequest) (command.UpdatePublishedCommand, error) {
	cmd := command.UpdatePublishedCommand{}

	// 除了ID，其他字段都是可选的，所以需要判断是否为空，如果为空则不赋值，否则赋值
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return command.UpdatePublishedCommand{}, err
	}
	cmd.ID = id

	if req.Title != nil {
		cmd.Title = *req.Title
	}

	if req.Content != nil {
		cmd.Content = *req.Content
	}

	if req.Abstract != nil {
		cmd.Abstract = *req.Abstract
	}

	if req.CoverUrl != nil {
		cmd.CoverUrl = *req.CoverUrl
	}

	// 处理可选的CategoryID
	if req.CategoryID != nil {
		categoryID, err := strconv.Atoi(*req.CategoryID)
		if err != nil {
			return command.UpdatePublishedCommand{}, err
		}
		cmd.CategoryID = categoryID
	}

	// 处理可选的ColumnID
	if req.ColumnID != nil {
		columnID, err := strconv.Atoi(*req.ColumnID)
		if err != nil {
			return command.UpdatePublishedCommand{}, err
		}
		cmd.ColumnID = columnID
	}

	// 处理可选的TagIDs
	if req.TagIDs != nil {
		var tagIDs []int
		for _, tagIDStr := range *req.TagIDs {
			tagID, err := strconv.Atoi(tagIDStr)
			if err != nil {
				return command.UpdatePublishedCommand{}, err
			}
			tagIDs = append(tagIDs, tagID)
		}
		cmd.TagIDs = tagIDs
	}

	return cmd, nil
}

func ConvertCreateDraftRequestToCommand(req request.CreateDraftRequest) (command.CreateDraftCommand, error) {
	cmd := command.CreateDraftCommand{
		Title:   req.Title,
		Content: req.Content,
	}

	if req.Abstract != nil {
		cmd.Abstract = *req.Abstract
	}

	if req.CoverUrl != nil {
		cmd.CoverUrl = *req.CoverUrl
	}

	// 处理可选的CategoryID
	if req.CategoryID != nil {
		categoryID, err := strconv.Atoi(*req.CategoryID)
		if err != nil {
			return command.CreateDraftCommand{}, err
		}
		cmd.CategoryID = categoryID
	}

	// 处理可选的ColumnID
	if req.ColumnID != nil {
		columnID, err := strconv.Atoi(*req.ColumnID)
		if err != nil {
			return command.CreateDraftCommand{}, err
		}
		cmd.ColumnID = columnID
	}

	// 转换TagIDs从[]string到[]int
	var tagIDs []int
	for _, tagIDStr := range req.TagIDs {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			return command.CreateDraftCommand{}, err
		}
		tagIDs = append(tagIDs, tagID)
	}
	cmd.TagIDs = tagIDs

	return cmd, nil
}

func ConvertUpdateDraftRequestToCommand(req request.UpdateDraftRequest) (command.UpdateDraftCommand, error) {
	cmd := command.UpdateDraftCommand{}

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return command.UpdateDraftCommand{}, err
	}
	cmd.ID = id

	if req.Title != nil {
		cmd.Title = *req.Title
	}

	if req.Content != nil {
		cmd.Content = *req.Content
	}

	if req.Abstract != nil {
		cmd.Abstract = *req.Abstract
	}

	if req.CoverUrl != nil {
		cmd.CoverUrl = *req.CoverUrl
	}

	// 处理可选的CategoryID
	if req.CategoryID != nil {
		categoryID, err := strconv.Atoi(*req.CategoryID)
		if err != nil {
			return command.UpdateDraftCommand{}, err
		}
		cmd.CategoryID = categoryID
	}

	// 处理可选的ColumnID
	if req.ColumnID != nil {
		columnID, err := strconv.Atoi(*req.ColumnID)
		if err != nil {
			return command.UpdateDraftCommand{}, err
		}
		cmd.ColumnID = columnID
	}

	// 转换TagIDs从[]string到[]int
	var tagIDs []int
	for _, tagIDStr := range req.TagIDs {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			return command.UpdateDraftCommand{}, err
		}
		tagIDs = append(tagIDs, tagID)
	}
	cmd.TagIDs = tagIDs

	return cmd, nil

}

func ConvertPublishDraftRequestToCommand(req request.PublishDraftRequest) (command.PublishDraftCommand, error) {
	cmd := command.PublishDraftCommand{
		Title:    req.Title,
		Abstract: req.Abstract,
		Content:  req.Content,
	}
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return command.PublishDraftCommand{}, err
	}
	cmd.ID = id

	// 处理可选字段CoverUrl
	if req.CoverUrl != nil {
		cmd.CoverUrl = *req.CoverUrl
	}

	// 转换CategoryID从string到int
	categoryID, err := strconv.Atoi(req.CategoryID)
	if err != nil {
		return command.PublishDraftCommand{}, err
	}
	cmd.CategoryID = categoryID

	// 转换ColumnID从string到int（可选字段）
	if req.ColumnID != nil {
		columnID, err := strconv.Atoi(*req.ColumnID)
		if err != nil {
			return command.PublishDraftCommand{}, err
		}
		cmd.ColumnID = columnID
	}

	// 转换TagIDs从[]string到[]int
	var tagIDs []int
	for _, tagIDStr := range req.TagIDs {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			return command.PublishDraftCommand{}, err
		}
		tagIDs = append(tagIDs, tagID)
	}
	cmd.TagIDs = tagIDs

	return cmd, nil
}

func ConvertDeleteArticleRequestToCommand(req request.DeleteArticleRequest) (command.DeleteArticleCommand, error) {
	return command.DeleteArticleCommand{
		ID: req.ID,
	}, nil
}

func ConvertUpdateArticleTopStatusRequestToCommand(req request.UpdateArticleTopStatusRequest) (command.UpdateArticleTopStatusCommand, error) {
	return command.UpdateArticleTopStatusCommand{
		ID:    req.ID,
		IsTop: req.IsTop,
	}, nil
}

func ConvertGetArticleDetailRequestToQuery(req request.GetArticleDetailRequest) (query.GetArticleDetailQuery, error) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return query.GetArticleDetailQuery{}, err
	}
	return query.GetArticleDetailQuery{
		ID: id,
	}, nil
}

func ConvertGetAllArticleListRequestToQuery(req request.GetAllArticleListRequest) (query.GetAllArticleListQuery, error) {
	var keyword string
	if req.Keyword == nil {
		keyword = ""
	}

	return query.GetAllArticleListQuery{
		Status:    req.Status,
		Page:      req.Page,
		PageSize:  req.PageSize,
		Keyword:   keyword,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}, nil
}

func ConvertGetAggregatedArticleListRequestToQuery(req request.GetAggregatedArticleListRequest) (query.GetAggregatedArticleListQuery, error) {

	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return query.GetAggregatedArticleListQuery{}, err
	}

	return query.GetAggregatedArticleListQuery{
		Type:     req.Type,
		ID:       id,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func ConvertGetArticleListByCategoryRequestToQuery(req request.GetArticleListByCategoryRequest) (query.GetArticleListByCategoryQuery, error) {
	id, err := strconv.Atoi(req.CategoryID)
	if err != nil {
		return query.GetArticleListByCategoryQuery{}, err
	}
	return query.GetArticleListByCategoryQuery{
		CategoryID: id,
		Page:       req.Page,
		PageSize:   req.PageSize,
		SortBy:     req.SortBy,
		SortOrder:  req.SortOrder,
	}, nil
}

func ConvertGetArticleListByColumnRequestToQuery(req request.GetArticleListByColumnRequest) (query.GetArticleListByColumnQuery, error) {
	id, err := strconv.Atoi(req.ColumnID)
	if err != nil {
		return query.GetArticleListByColumnQuery{}, err
	}
	return query.GetArticleListByColumnQuery{
		ColumnID:  id,
		Page:      req.Page,
		PageSize:  req.PageSize,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}, nil
}

func ConvertGetArticleListByTagRequestToQuery(req request.GetArticleListByTagRequest) (query.GetArticleListByTagQuery, error) {
	id, err := strconv.Atoi(req.TagID)
	if err != nil {
		return query.GetArticleListByTagQuery{}, err
	}
	return query.GetArticleListByTagQuery{
		TagID:     id,
		Page:      req.Page,
		PageSize:  req.PageSize,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}, nil
}

func ConvertGetArticleListByKeywordRequestToQuery(req request.GetArticleListByKeywordRequest) (query.GetArticleListByKeywordQuery, error) {
	return query.GetArticleListByKeywordQuery{
		Keyword:   req.Keyword,
		Page:      req.Page,
		PageSize:  req.PageSize,
		SearchIn:  req.SearchIn,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}, nil
}

// ================================== 后端模型转换成前端响应 ==================================
func ConvertArticleDetailResultToResponse(res *result.ArticleDetailResult) *response.GetArticleDetailResponse {
	var categoryResp *response.CategoryResp
	if res.Category != nil {
		categoryResp = &response.CategoryResp{
			CategoryId:   res.Category.ID,   // 字段名映射：ID → categoryId
			CategoryName: res.Category.Name, // Name → categoryName
		}
	}

	var columnResp *response.ColumnResp
	if res.Column != nil {
		columnResp = &response.ColumnResp{
			ColumnId:   res.Column.ID,
			ColumnName: res.Column.Name,
		}
	}

	tagResps := make([]*response.TagResp, 0, len(res.Tags))
	for _, tagResult := range res.Tags {
		tagResps = append(tagResps, &response.TagResp{
			TagId:   tagResult.ID,
			TagName: tagResult.Name,
		})
	}

	return &response.GetArticleDetailResponse{
		ID:          res.ID,
		Title:       res.Title,
		Content:     res.Content,
		Abstract:    res.Abstract,
		CoverUrl:    res.CoverUrl,
		Status:      res.Status,
		IsTop:       res.IsTop,
		ViewCount:   res.ViewCount,
		LikeCount:   res.LikeCount,
		CreateTime:  res.CreateTime,
		PublishTime: res.PublishTime,
		UpdatedAt:   res.UpdateTime,
		Category:    categoryResp,
		Column:      columnResp,
		Tags:        tagResps,
	}
}

func ConvertAllArticleListResultToResponse(res *result.AllArticleListResult) *response.GetAllArticleListResponse {
	if res == nil {
		return &response.GetAllArticleListResponse{
			Items:      []*response.ArticleListItemResp{}, // 空列表返回空数组
			Total:      0,
			TotalPages: 0,
			Page:       0,
			PageSize:   0,
		}
	}

	// 转换列表项
	// items := make([]*response.ArticleListItemResp, 0, len(res.Items))
	// for _, itemResult := range res.Items {
	// 	// 转换分类
	// 	var categoryResp *response.CategoryResp
	// 	if itemResult.Category != nil {
	// 		categoryResp = &response.CategoryResp{
	// 			CategoryId:   itemResult.Category.ID,
	// 			CategoryName: itemResult.Category.Name,
	// 		}
	// 	}

	// 	// 转换专栏
	// 	var columnResp *response.ColumnResp
	// 	if itemResult.Column != nil {
	// 		columnResp = &response.ColumnResp{
	// 			ColumnId:   itemResult.Column.ID,
	// 			ColumnName: itemResult.Column.Name,
	// 		}
	// 	}

	// 	// 转换标签
	// 	tagResps := make([]*response.TagResp, 0, len(itemResult.Tags))
	// 	for _, tagResult := range itemResult.Tags {
	// 		tagResps = append(tagResps, &response.TagResp{
	// 			TagId:   tagResult.ID,
	// 			TagName: tagResult.Name,
	// 		})
	// 	}

	// 	// 组装单条列表项
	// 	items = append(items, &response.ArticleListItemResp{
	// 		ID:          itemResult.ID,
	// 		Title:       itemResult.Title,
	// 		Abstract:    itemResult.Abstract,
	// 		Category:    categoryResp,
	// 		Column:      columnResp,
	// 		Tags:        tagResps,
	// 		ViewCount:   itemResult.ViewCount,
	// 		PublishTime: itemResult.PublishTime,
	// 		IsTop:       itemResult.IsTop,
	// 	})
	// }

	items := converArticleListItemResultToResponse(res.Items)
	// 组装整体响应
	return &response.GetAllArticleListResponse{
		Items:      items,
		Total:      res.Total,
		Page:       res.Page,
		PageSize:   res.PageSize,
		TotalPages: res.TotalPages,
	}
}

func ConvertHomePageArticleListResultToResponse(res *result.HomePageArticleListResult) *response.GetHomePageArticleListResponse {
	if res == nil {
		return &response.GetHomePageArticleListResponse{
			Items:     []*response.ArticleListItemResp{}, // 空列表返回空数组
			Type:      "",
			SortOrder: "",
			SortBy:    "",
		}
	}

	items := converArticleListItemResultToResponse(res.Items)

	return &response.GetHomePageArticleListResponse{
		Items:     items,
		Type:      res.Type,
		SortOrder: res.SortOrder,
		SortBy:    res.SortBy,
	}
}

func ConvertAggregatedArticleListResultToResponse(res *result.AggregatedArticleListResult) *response.GetAggregatedArticleListResponse {
	if res == nil {
		return &response.GetAggregatedArticleListResponse{
			Items:      []*response.ArticleListItemResp{}, // 空列表返回空数组
			Type:       "",
			Total:      0,
			Page:       0,
			PageSize:   0,
			TotalPages: 0,
		}
	}

	// 转换列表项
	items := converArticleListItemResultToResponse(res.Items)

	return &response.GetAggregatedArticleListResponse{
		Items:      items,
		Type:       res.Type,
		Total:      res.Total,
		Page:       res.Page,
		PageSize:   res.PageSize,
		TotalPages: res.TotalPages,
	}
}

func converArticleListItemResultToResponse(itemsResult []*result.ArticleListItemResult) []*response.ArticleListItemResp {
	items := make([]*response.ArticleListItemResp, 0, len(itemsResult))
	for _, itemResult := range itemsResult {
		// 转换分类
		var categoryResp *response.CategoryResp
		if itemResult.Category != nil {
			categoryResp = &response.CategoryResp{
				CategoryId:   itemResult.Category.ID,
				CategoryName: itemResult.Category.Name,
			}
		}

		// 转换专栏
		var columnResp *response.ColumnResp
		if itemResult.Column != nil {
			columnResp = &response.ColumnResp{
				ColumnId:   itemResult.Column.ID,
				ColumnName: itemResult.Column.Name,
			}
		}

		// 转换标签
		tagResps := make([]*response.TagResp, 0, len(itemResult.Tags))
		for _, tagResult := range itemResult.Tags {
			tagResps = append(tagResps, &response.TagResp{
				TagId:   tagResult.ID,
				TagName: tagResult.Name,
			})
		}

		// 组装单条列表项
		items = append(items, &response.ArticleListItemResp{
			ID:          itemResult.ID,
			Title:       itemResult.Title,
			Abstract:    itemResult.Abstract,
			CoverUrl:    itemResult.CoverUrl,
			Category:    categoryResp,
			Column:      columnResp,
			Tags:        tagResps,
			ViewCount:   itemResult.ViewCount,
			CreateTime:  itemResult.CreateTime,
			PublishTime: itemResult.PublishTime,
			IsTop:       itemResult.IsTop,
		})
	}
	return items
}
