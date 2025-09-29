package converter

import (
	"errors"
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/dto/request"
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/dto/response"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/command"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/query"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/result"
	"log"
)

func ConvertCreateTagRequestToCommand(req request.CreateTagRequest) (command.CreateTagCommand, error) {
	cmd := command.CreateTagCommand{
		Name:        req.Name,
		Description: req.Description,
	}
	return cmd, nil
}

func ConvertUpdateTagRequestToCommand(req request.UpdateTagRequest) (command.UpdateTagCommand, error) {
	cmd := command.UpdateTagCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}
	return cmd, nil
}

func ConvertDeleteTagRequestToCommand(req request.DeleteTagRequest) (command.DeleteTagCommand, error) {
	cmd := command.DeleteTagCommand{
		ID: req.ID,
	}
	return cmd, nil
}

func ConvertGetTagListRequestToQuery(req request.GetTagListRequest) (query.GetTagListQuery, error) {
	qry := query.GetTagListQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return qry, nil
}

func ConvertGetTagDetailRequestToQuery(req request.GetTagDetailRequest) (query.GetTagDetailQuery, error) {
	qry := query.GetTagDetailQuery{
		ID: req.ID,
	}
	return qry, nil
}

// ================================================================

func ConvertTagDetailResultToResponse(res *result.TagDetailResult) (response.TagDetailResponse, error) {
	if res == nil {
		return response.TagDetailResponse{}, errors.New("nil result")
	}

	return response.TagDetailResponse{
		ID:           res.ID,
		Name:         res.Name,
		Description:  res.Description,
		CreateTime:   res.CreateTime,
		UpdateTime:   res.UpdateTime,
		ArticleCount: res.ArticleCount,
	}, nil
}

func ConvertTagListResultToResponse(res *result.TagListResult) (response.TagListResponse, error) {
	if res == nil {
		return response.TagListResponse{}, errors.New("nil result")
	}

	items := make([]*response.TagDetailResponse, 0, len(res.Items))
	for _, item := range res.Items {
		resp, err := ConvertTagDetailResultToResponse(item)
		if err != nil {
			// 记录错误但继续处理其他项
			log.Printf("ConvertTagDetailResultToResponse error: %v", err)
			continue
		}
		items = append(items, &resp)
	}

	return response.TagListResponse{
		Items: items,
		Total: res.Total,
	}, nil
}
