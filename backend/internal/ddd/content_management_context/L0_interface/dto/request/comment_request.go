package request

type CommentSubmitReq struct {
	ParentID  *uint  `json:"parentId,omitempty"`
	ArticleID int    `json:"articleId" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
