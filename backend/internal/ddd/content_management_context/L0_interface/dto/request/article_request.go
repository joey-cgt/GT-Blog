package request

// 发布新文章的请求参数
type CreateAndPublishRequest struct {
	Title      string   `json:"title" validate:"required,min=1,max=50"`     // 文章标题（必填，1-50字符）
	Content    string   `json:"content" validate:"required"`                // 文章内容（必填）
	Abstract   string   `json:"abstract" validate:"required,min=1,max=100"` // 文章摘要（必填，1-100字符）
	CoverUrl   *string  `json:"coverUrl" validate:"omitempty,url"`          // 封面图URL（可选，需合法URL格式）
	CategoryID string   `json:"categoryId" validate:"required"`             // 分类ID（必填）
	ColumnID   *string  `json:"columnId" validate:"omitempty"`              // 专栏ID（可选）
	TagIDs     []string `json:"tagIds" validate:"omitempty,dive,required"`  // 标签ID列表（可选，每个ID必填）
	Status     int      `json:"status" validate:"required"`                 // 文章状态（必填）
}

// 更新已发布文章的请求参数（支持部分字段更新）
type UpdatePublishedRequest struct {
	ID         string    `json:"id" validate:"required"`
	Title      *string   `json:"title" validate:"omitempty,min=1,max=50"` // 可选更新字段，不为空时校验
	Content    *string   `json:"content" validate:"omitempty"`
	Abstract   *string   `json:"abstract" validate:"omitempty,min=1,max=100"`
	CoverUrl   *string   `json:"coverUrl" validate:"omitempty,url"`
	CategoryID *string   `json:"categoryId" validate:"omitempty"`
	ColumnID   *string   `json:"columnId" validate:"omitempty"`
	TagIDs     *[]string `json:"tagIds" validate:"omitempty,dive,required"`
	Status     *int      `json:"status" validate:"omitempty"`
}

// 保存草稿的请求参数
type CreateDraftRequest struct {
	Title      string   `json:"title" validate:"required,min=1,max=50"`
	Content    string   `json:"content" validate:"required"`
	Abstract   *string  `json:"abstract" validate:"required,min=1,max=100"`
	CoverUrl   *string  `json:"coverUrl" validate:"omitempty,url"`
	CategoryID *string  `json:"categoryId" validate:"required"`
	ColumnID   *string  `json:"columnId" validate:"omitempty"`
	TagIDs     []string `json:"tagIds" validate:"omitempty,dive,required"`
}

// 更新草稿的请求参数
type UpdateDraftRequest struct {
	ID         string   `json:"id" validate:"required"`
	Title      *string  `json:"title" validate:"omitempty,min=1,max=50"`
	Content    *string  `json:"content" validate:"omitempty"`
	Abstract   *string  `json:"abstract" validate:"omitempty,min=1,max=100"`
	CoverUrl   *string  `json:"coverUrl" validate:"omitempty,url"`
	CategoryID *string  `json:"categoryId" validate:"omitempty"`
	ColumnID   *string  `json:"columnId" validate:"omitempty"`
	TagIDs     []string `json:"tagIds" validate:"omitempty,dive,required"`
}

// 发布草稿的请求参数
type PublishDraftRequest struct {
	ID         string   `json:"id" validate:"required"`
	Title      string   `json:"title" validate:"required,min=1,max=50"`     // 文章标题（必填，1-50字符）
	Content    string   `json:"content" validate:"required"`                // 文章内容（必填）
	Abstract   string   `json:"abstract" validate:"required,min=1,max=100"` // 文章摘要（必填，1-100字符）
	CoverUrl   *string  `json:"coverUrl" validate:"omitempty,url"`          // 封面图URL（可选，需合法URL格式）
	CategoryID string   `json:"categoryId" validate:"required"`             // 分类ID（必填）
	ColumnID   *string  `json:"columnId" validate:"omitempty"`              // 专栏ID（可选）
	TagIDs     []string `json:"tagIds" validate:"omitempty,dive,required"`  // 标签ID列表（可选，每个ID必填）
}

// 删除文章（已发布/草稿）的请求参数
type DeleteArticleRequest struct {
	ID int `uri:"id" json:"id" validate:"required"` // 草稿ID（必填）
}

// 置顶/取消置顶的请求参数
type UpdateArticleTopStatusRequest struct {
	ID    int  `json:"id" validate:"required"`             // 文章ID（必填）
	IsTop bool `json:"is_top" validate:"required,boolean"` // 必须明确是否置顶
}

// 获取文章详情的请求参数
type GetArticleDetailRequest struct {
	ID string `uri:"id" validate:"required"` // 文章ID（必填）
}

// 获取全部文章的请求参数
type GetAllArticleListRequest struct {
	Status    int     `form:"status" validate:"omitempty,oneof=0 1"`
	Page      int     `form:"page" validate:"omitempty,min=1"`                                                 // 页码，默认1
	PageSize  int     `form:"pageSize" validate:"omitempty,min=1,max=100"`                                     // 每页数量，默认10，最大100
	Keyword   *string `form:"keyword" validate:"omitempty"`                                                    // 关键词搜索（标题/内容）
	SortBy    string  `form:"sortBy" validate:"omitempty,oneof=created_at published_at view_count like_count"` // 排序字段
	SortOrder string  `form:"sortOrder" validate:"omitempty,oneof=asc desc"`                                   // 排序方向
}

// 获取聚合（专栏/分类/标签）文章列表的请求参数
type GetAggregatedArticleListRequest struct {
	Type     string `form:"type" validate:"required,oneof=column tag category"`
	ID       string `form:"id" validate:"required"`
	Page     int    `form:"page" validate:"omitempty,min=1"`
	PageSize int    `form:"pageSize" validate:"omitempty,min=1,max=100"`
}

// 获取某分类下的文章列表的请求参数
type GetArticleListByCategoryRequest struct {
	CategoryID string `form:"categoryId" validate:"required"`                                                  // 分类ID（必填）
	Page       int    `form:"page" validate:"omitempty,min=1"`                                                 // 页码，默认1
	PageSize   int    `form:"pageSize" validate:"omitempty,min=1,max=100"`                                     // 每页数量，默认10，最大100
	SortBy     string `form:"sortBy" validate:"omitempty,oneof=created_at published_at view_count like_count"` // 排序字段
	SortOrder  string `form:"sortOrder" validate:"omitempty,oneof=asc desc"`                                   // 排序方向
}

// 获取某专栏下的文章列表的请求参数
type GetArticleListByColumnRequest struct {
	ColumnID  string `form:"columnId" validate:"required"`                                                    // 专栏ID（必填）
	Page      int    `form:"page" validate:"omitempty,min=1"`                                                 // 页码，默认1
	PageSize  int    `form:"pageSize" validate:"omitempty,min=1,max=100"`                                     // 每页数量，默认10，最大100
	SortBy    string `form:"sortBy" validate:"omitempty,oneof=created_at published_at view_count like_count"` // 排序字段
	SortOrder string `form:"sortOrder" validate:"omitempty,oneof=asc desc"`                                   // 排序方向
}

// 获取某标签下的文章列表的请求参数
type GetArticleListByTagRequest struct {
	TagID     string `form:"tagId" validate:"required"`                                                       // 标签ID（必填）
	Page      int    `form:"page" validate:"omitempty,min=1"`                                                 // 页码，默认1
	PageSize  int    `form:"pageSize" validate:"omitempty,min=1,max=100"`                                     // 每页数量，默认10，最大100
	SortBy    string `form:"sortBy" validate:"omitempty,oneof=created_at published_at view_count like_count"` // 排序字段
	SortOrder string `form:"sortOrder" validate:"omitempty,oneof=asc desc"`                                   // 排序方向
}

// 获取关键字搜索文章列表的请求参数
type GetArticleListByKeywordRequest struct {
	Keyword   string `form:"keyword" validate:"required,min=1"`                                   // 搜索关键词（必填）
	Page      int    `form:"page" validate:"omitempty,min=1"`                                     // 页码，默认1
	PageSize  int    `form:"pageSize" validate:"omitempty,min=1,max=100"`                         // 每页数量，默认10，最大100
	SearchIn  string `form:"searchIn" validate:"omitempty,oneof=title content both"`              // 搜索范围：title(标题),content(内容),both(两者都搜)
	SortBy    string `form:"sortBy" validate:"omitempty,oneof=created_at published_at relevance"` // 排序字段
	SortOrder string `form:"sortOrder" validate:"omitempty,oneof=asc desc"`                       // 排序方向
}
