package result

type ColumnDetailResult struct {
	ID           int
	Name         string
	Description  string
	CoverUrl     string
	CreateTime   string
	UpdateTime   string
	ArticleCount int
	ViewCount    int
	LikeCount    int
}

type ColumnListResult struct {
	Total     int
	Page      int
	PageSize  int
	PageCount int
	List      []*ColumnDetailResult
}
