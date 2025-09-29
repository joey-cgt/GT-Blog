package result

type CategoryDetailResult struct {
	ID           int
	Name         string
	Description  string
	CreateTime   string
	UpdateTime   string
	ArticleCount int
}

type CategoryListResult struct {
	Total int
	List  []*CategoryDetailResult
}
