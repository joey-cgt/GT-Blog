package query

// 获取文章详情命令
type GetArticleDetailQuery struct {
	ID int // 文章ID
}

// 获取全部文章命令
type GetAllArticleListQuery struct {
	Status    int    // 0 草稿 1 已发布
	Page      int    // 页码
	PageSize  int    // 每页数量
	Keyword   string // 关键词搜索
	SortBy    string // 排序字段
	SortOrder string // 排序方向
}

// 获取主页（置顶/最新/热门）文章命令
type GetHomePageArticleListQuery struct {
	Type      string // 类型：top(置顶)、new(最新)、hot(热门)
	SortBy    string // 排序字段
	SortOrder string // 排序方向
}

type GetAggregatedArticleListQuery struct {
	Type     string // 聚合类型：column(专栏)、tag(标签)、category(分类)
	ID       int    // ID
	Page     int    // 页码
	PageSize int    // 每页数量
}

// 获取某分类下的文章列表命令
type GetArticleListByCategoryQuery struct {
	CategoryID int    // 分类ID
	Page       int    // 页码
	PageSize   int    // 每页数量
	SortBy     string // 排序字段
	SortOrder  string // 排序方向
}

// 获取某专栏下的文章列表命令
type GetArticleListByColumnQuery struct {
	ColumnID  int    // 专栏ID
	Page      int    // 页码
	PageSize  int    // 每页数量
	SortBy    string // 排序字段
	SortOrder string // 排序方向
}

// 获取某标签下的文章列表命令
type GetArticleListByTagQuery struct {
	TagID     int    // 标签ID
	Page      int    // 页码
	PageSize  int    // 每页数量
	SortBy    string // 排序字段
	SortOrder string // 排序方向
}

// 获取关键字搜索的文章列表命令
type GetArticleListByKeywordQuery struct {
	Keyword   string // 搜索关键词
	Page      int    // 页码
	PageSize  int    // 每页数量
	SearchIn  string // 搜索范围，可选值："title"(仅标题)、"content"(仅内容)、"both"(标题和内容都搜索)
	SortBy    string // 排序字段
	SortOrder string // 排序方向
}
