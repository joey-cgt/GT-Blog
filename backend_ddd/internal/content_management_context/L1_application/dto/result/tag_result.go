package result

type TagDetailResult struct {
	ID           int
	Name         string
	Description  string
	CreateTime   string
	UpdateTime   string
	ArticleCount int
}

type TagListResult struct {
	Items     []*TagDetailResult
	Total     int
}
