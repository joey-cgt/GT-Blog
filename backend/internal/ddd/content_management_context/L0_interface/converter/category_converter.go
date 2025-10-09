package converter

import (
	"errors"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/response"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/command"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/query"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"log"
)

func ConvertCreateCategoryRequestToCommand(req request.CreateCategoryRequest) (command.CreateCategoryCommand, error) {
	cmd := command.CreateCategoryCommand{
		Name:        req.Name,
		Description: req.Description,
	}
	return cmd, nil
}

func ConvertUpdateCategoryRequestToCommand(req request.UpdateCategoryRequest) (command.UpdateCategoryCommand, error) {
	cmd := command.UpdateCategoryCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}
	return cmd, nil
}

func ConvertDeleteCategoryRequestToCommand(req request.DeleteCategoryRequest) (command.DeleteCategoryCommand, error) {
	cmd := command.DeleteCategoryCommand{
		ID: req.ID,
	}
	return cmd, nil
}

func ConvertGetCategoryDetailRequestToQuery(req request.GetCategoryDetailRequest) (query.GetCategoryDetailQuery, error) {
	qry := query.GetCategoryDetailQuery{
		ID: req.ID,
	}
	return qry, nil
}

func ConvertGetCategoryListRequestToQuery(req request.GetCategoryListRequest) (query.GetCategoryListQuery, error) {
	qry := query.GetCategoryListQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return qry, nil
}

// ============================================================================================

func ConvertCategoryDetailResultToResponse(res *result.CategoryDetailResult) (response.GetCategoryDetailResponse, error) {
	resp := response.GetCategoryDetailResponse{
		ID:           res.ID,
		Name:         res.Name,
		Description:  res.Description,
		ArticleCount: res.ArticleCount,
		CreateTime:   res.CreateTime,
		UpdateTime:   res.UpdateTime,
	}
	return resp, nil
}

func ConvertCategoryListResultToResponse(res *result.CategoryListResult) (response.GetCategoryListResponse, error) {
	if res == nil {
		return response.GetCategoryListResponse{}, errors.New("nil result")
	}

	items := make([]*response.GetCategoryDetailResponse, 0, len(res.List))
	for _, item := range res.List {
		resp, err := ConvertCategoryDetailResultToResponse(item)
		if err != nil {
			// 记录错误但继续处理其他项
			log.Printf("ConvertCategoryDetailResultToResponse error: %v", err)
			continue
		}
		items = append(items, &resp)
	}

	return response.GetCategoryListResponse{
		Total: res.Total,
		Items: items,
	}, nil
}
