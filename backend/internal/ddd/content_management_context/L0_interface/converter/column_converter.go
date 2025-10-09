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

func ConvertCreateColumnRequestToCommand(req request.CreateColumnRequest) (command.CreateColumnCommand, error) {
	cmd := command.CreateColumnCommand{
		Name:        req.Name,
		Description: req.Description,
		CoverUrl:    req.CoverUrl,
	}

	return cmd, nil
}

func ConvertUpdateColumnRequestToCommand(req request.UpdateColumnRequest) (command.UpdateColumnCommand, error) {
	cmd := command.UpdateColumnCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		CoverUrl:    req.CoverUrl,
	}
	return cmd, nil
}

func ConvertDeleteColumnRequestToCommand(req request.DeleteColumnRequest) (command.DeleteColumnCommand, error) {
	cmd := command.DeleteColumnCommand{
		ID: req.ID,
	}
	return cmd, nil
}

func ConvertGetColumnDetailRequestToQuery(req request.GetColumnDetailRequest) (query.GetColumnDetailQuery, error) {
	cmd := query.GetColumnDetailQuery{
		ID: req.ID,
	}
	return cmd, nil
}

func ConvertGetColumnListRequestToQuery(req request.GetColumnListRequest) (query.GetColumnListQuery, error) {
	qry := query.GetColumnListQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return qry, nil
}

// =======================================================================================

func ConvertColumnListResultToResponse(res *result.ColumnListResult) (response.GetColumnListResponse, error) {
	if res == nil {
		return response.GetColumnListResponse{}, errors.New("nil result")
	}

	items := make([]*response.GetColumnDetailResponse, 0, len(res.List))
	for _, item := range res.List {
		resp, err := ConvertColumnDetailResultToResponse(item)
		if err != nil {
			// 记录错误但继续处理其他项
			log.Printf("ConvertColumnDetailResultToResponse error: %v", err)
			continue
		}
		items = append(items, &resp)
	}

	return response.GetColumnListResponse{
		Items:      items,
		Total:      res.Total,
		Page:       res.Page,
		PageSize:   res.PageSize,
		TotalPages: (res.Total + res.PageSize - 1) / res.PageSize, // 计算总页数
	}, nil
}

func ConvertColumnDetailResultToResponse(res *result.ColumnDetailResult) (response.GetColumnDetailResponse, error) {
	resp := response.GetColumnDetailResponse{
		ID:           res.ID,
		Name:         res.Name,
		Description:  res.Description,
		CoverUrl:     res.CoverUrl,
		CreateTime:   res.CreateTime,
		UpdateTime:   res.UpdateTime,
		ArticleCount: res.ArticleCount,
		ViewCount:    res.ViewCount,
		LikeCount:    res.LikeCount,
	}
	return resp, nil
}
