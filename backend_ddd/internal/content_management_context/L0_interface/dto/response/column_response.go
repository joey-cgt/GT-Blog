package response

type GetColumnDetailResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CoverUrl     string `json:"coverUrl"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	ArticleCount int    `json:"articleCount"`
	ViewCount    int    `json:"viewCount"`
	LikeCount    int    `json:"likeCount"`
	IsTop        bool   `json:"isTop"`
}

type GetColumnListResponse struct {
	Items      []*GetColumnDetailResponse `json:"items"`
	Total      int                        `json:"total"`
	Page       int                        `json:"page"`
	PageSize   int                        `json:"pageSize"`
	TotalPages int                        `json:"totalPages"`
}
