package result

type CategoryResult struct {
	ID   int
	Name string
}

type ColumnResult struct {
	ID   int
	Name string
}

type TagResult struct {
	ID   int
	Name string
}

type ArticleDetailResult struct {
	ID          int
	Title       string
	Abstract    string
	Content     string
	Column      *ColumnResult
	Category    *CategoryResult
	Tags        []*TagResult
	CoverUrl    string
	Status      int
	IsTop       bool
	ViewCount   int
	LikeCount   int
	CreateTime  string
	PublishTime string
	UpdateTime  string
}

type ArticleListItemResult struct {
	ID           int             // 文章ID
	Title        string          // 标题
	Abstract     string          // 摘要（列表页无需完整内容）
	Category     *CategoryResult // 分类信息（ID+Name）
	Column       *ColumnResult   // 专栏信息（ID+Name）
	Tags         []*TagResult    // 标签列表（ID+Name）
	ViewCount    int             // 浏览量
	LikeCount    int             // 点赞量
	CommentCount int             // 评论量
	CreateTime   string          // 创建时间（格式化后）
	PublishTime  string          // 发布时间（格式化后）
	IsTop        bool            // 是否置顶
}

type AllArticleListResult struct {
	Items      []*ArticleListItemResult // 文章列表项
	Total      int                      // 总条数（用于分页）
	Page       int                      // 当前页码
	PageSize   int                      // 每页条数
	TotalPages int                      // 总页数（方便前端计算）
}

type HomePageArticleListResult struct {
	Items     []*ArticleListItemResult //
	Type      string
	SortBy    string
	SortOrder string
}

type AggregatedArticleListResult struct {
	Items      []*ArticleListItemResult
	Type       string
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}
