package response

type CategoryResp struct {
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// ColumnResp 专栏关联数据的响应DTO
type ColumnResp struct {
	ColumnId   int    `json:"columnId"`
	ColumnName string `json:"columnName"`
}

// TagResp 标签关联数据的响应DTO
type TagResp struct {
	TagId   int    `json:"tagId"`
	TagName string `json:"tagName"`
}

// 获取全部文章的响应
type GetAllArticleListResponse struct {
	Items      []*ArticleListItemResp `json:"items"`      // 文章列表数据S
	Total      int                    `json:"total"`      // 总记录数
	Page       int                    `json:"page"`       // 当前页码
	PageSize   int                    `json:"pageSize"`   // 每页数量
	TotalPages int                    `json:"totalPages"` // 总页数
}

// 文章列表项
type ArticleListItemResp struct {
	ID          int           `json:"id"`          // 文章ID
	Title       string        `json:"title"`       // 文章标题
	Abstract    string        `json:"abstract"`    // 文章摘要
	CoverUrl    string        `json:"coverUrl"`    // 封面图URL
	Status      int           `json:"status"`      // 文章状态
	IsTop       bool          `json:"isTop"`       // 是否置顶
	ViewCount   int           `json:"viewCount"`   // 浏览次数
	LikeCount   int           `json:"likeCount"`   // 点赞次数
	CreateTime  string        `json:"createTime"`  // 创建时间
	PublishTime string        `json:"publishTime"` // 发布时间
	Category    *CategoryResp `json:"category"`    // 分类信息
	Column      *ColumnResp   `json:"column"`      // 专栏信息
	Tags        []*TagResp    `json:"tags"`        // 标签列表
}

// 获取单篇文章内容的响应
type GetArticleDetailResponse struct {
	ID           int           `json:"id"`           // 文章ID
	Title        string        `json:"title"`        // 文章标题
	Content      string        `json:"content"`      // 文章内容
	Abstract     string        `json:"abstract"`     // 文章摘要
	CoverUrl     string        `json:"coverUrl"`     // 封面图URL
	Status       int           `json:"status"`       // 文章状态
	IsTop        bool          `json:"isTop"`        // 是否置顶
	ViewCount    int           `json:"viewCount"`    // 浏览次数
	LikeCount    int           `json:"likeCount"`    // 点赞次数
	CommentCount int           `json:"commentCount"` // 评论数
	CreateTime   string        `json:"createTime"`   // 创建时间
	PublishTime  string        `json:"publishTime"`  // 发布时间
	UpdatedAt    string        `json:"updateTime"`   // 最后更新时间
	Category     *CategoryResp `json:"category"`     // 分类信息
	Column       *ColumnResp   `json:"column"`       // 专栏信息
	Tags         []*TagResp    `json:"tags"`         // 标签列表
}

type GetHomePageArticleListResponse struct {
	Items     []*ArticleListItemResp `json:"items"`
	Type      string                 `json:"type"`
	SortOrder string                 `json:"sortOrder"`
	SortBy    string                 `json:"sortBy"`
}

type GetAggregatedArticleListResponse struct {
	Items      []*ArticleListItemResp `json:"items"`
	Type       string                 `json:"type"`
	Total      int                    `json:"total"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"pageSize"`
	TotalPages int                    `json:"totalPages"`
}
