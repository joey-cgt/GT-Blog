package response

type GetCategoryDetailResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ArticleCount int    `json:"articleCount"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	ArticleCount int    `json:"articleCount"`
}

type GetCategoryListResponse struct {
	Items      []*GetCategoryDetailResponse `json:"items"`
	Total      int                          `json:"total"`
	Page       int                          `json:"page"`
	PageSize   int                          `json:"pageSize"`
	TotalPages int                          `json:"totalPages"`
}
