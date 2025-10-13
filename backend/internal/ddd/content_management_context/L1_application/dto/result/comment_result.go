package result

type CommentsResult struct {
	TotalCount int
	Comments   []CommentItemResult
	Page       int
	PageSize   int
}

type CommentItemResult struct {
	ID        uint
	Nickname  string
	Email     string
	Content   string
	ArticleID uint
	ParentID  *uint
	Children  []CommentItemResult
	CreatedAt string
}
