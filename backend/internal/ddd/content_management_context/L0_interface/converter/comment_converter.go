package converter

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/response"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
)

// ConvertGetArticleCommentsToResponse 将 CommentsResult 转换为 CommentResponse
func ConvertGetArticleCommentsToResponse(result result.CommentsResult) (response.CommentResponse, error) {
	// 创建响应对象
	resp := response.CommentResponse{
		TotalCount: result.TotalCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
	}

	// 转换评论列表（处理嵌套结构）
	resp.Comments = convertCommentItems(result.Comments)

	return resp, nil
}

// convertCommentItems 递归转换评论项列表，处理嵌套结构
func convertCommentItems(resultItems []result.CommentItemResult) []response.CommentItem {
	responseItems := make([]response.CommentItem, 0, len(resultItems))

	for _, resultItem := range resultItems {
		// 创建响应评论项
		responseItem := response.CommentItem{
			ID:        resultItem.ID,
			Username:  resultItem.Username,
			Email:     resultItem.Email,
			Content:   resultItem.Content,
			ArticleID: resultItem.ArticleID,
			ParentID:  resultItem.ParentID,
			CreatedAt: resultItem.CreatedAt,
		}

		// 递归转换子评论
		if len(resultItem.Children) > 0 {
			responseItem.Children = convertCommentItems(resultItem.Children)
		}

		responseItems = append(responseItems, responseItem)
	}

	return responseItems
}
