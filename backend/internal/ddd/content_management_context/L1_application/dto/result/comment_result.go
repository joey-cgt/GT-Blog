package result

type CommentsResult struct {
	TotalCount int                 `json:"total_count"`
	Comments   []CommentItemResult `json:"comments"`
	Page       int                 `json:"page,omitempty"`
	PageSize   int                 `json:"page_size,omitempty"`
}

type CommentItemResult struct {
	ID        uint                `json:"id"`
	Username  string              `json:"nickname"`
	Email     string              `json:"email,omitempty"` // 不返回给前端，仅用于后端处理
	Content   string              `json:"content"`
	ArticleID uint                `json:"article_id,omitempty"`
	ParentID  *uint               `json:"parent_id,omitempty"`
	Children  []CommentItemResult `json:"children,omitempty"`
	CreatedAt string              `json:"created_at"`
}
