package command

type CreateAndPublishCommand struct {
	Title      string
	Abstract   string
	Content    string
	ColumnID   int
	CategoryID int
	TagIDs     []int
	CoverUrl   string
}

type UpdatePublishedCommand struct {
	ID         int
	Title      string
	Abstract   string
	Content    string
	ColumnID   int
	CategoryID int
	TagIDs     []int
	CoverUrl   string
}

type CreateDraftCommand struct {
	Title      string
	Content    string
	Abstract   string
	CoverUrl   string
	CategoryID int
	ColumnID   int
	TagIDs     []int
}

type UpdateDraftCommand struct {
	ID         int
	Title      string
	Content    string
	Abstract   string
	CoverUrl   string
	CategoryID int
	ColumnID   int
	TagIDs     []int
}

type PublishDraftCommand struct {
	ID         int
	Title      string
	Abstract   string
	Content    string
	ColumnID   int
	CategoryID int
	TagIDs     []int
	CoverUrl   string
}

// 删除文章命令
type DeleteArticleCommand struct {
	ID int
}

// 更新文章置顶状态命令
type UpdateArticleTopStatusCommand struct {
	ID    int  // 文章ID
	IsTop bool // 是否置顶
}