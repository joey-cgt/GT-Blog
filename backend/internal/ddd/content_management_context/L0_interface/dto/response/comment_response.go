package response

type CommentResponse struct {
	TotalCount int           `json:"total_count"`
	Comments   []CommentItem `json:"comments"`
	Page       int           `json:"page,omitempty"`
	PageSize   int           `json:"page_size,omitempty"`
}

type CommentItem struct {
	ID        uint          `json:"id"`
	Username  string        `json:"nickname"`
	Email     string        `json:"email,omitempty"` // 不返回给前端，仅用于后端处理
	Content   string        `json:"content"`
	ArticleID uint          `json:"articleId,omitempty"`
	ParentID  *uint         `json:"parentId,omitempty"`
	Children  []CommentItem `json:"children,omitempty"`
	CreatedAt string        `json:"createdAt"`
}
