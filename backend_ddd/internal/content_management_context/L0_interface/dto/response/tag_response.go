package response

type CreateTagResponse struct {
	ID string `json:"id"`
}

type UpdateTagResponse struct {
	ID string `json:"id"`
}

type TagDetailResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	ArticleCount int    `json:"articleCount"`
}

type TagListResponse struct {
	Items      []*TagDetailResponse `json:"items"`
	Total      int                  `json:"total"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"pageSize"`
	TotalPages int                  `json:"totalPages"`
}
