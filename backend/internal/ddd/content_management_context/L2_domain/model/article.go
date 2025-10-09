package model

import (
	"fmt"
	"time"
)

type Article struct {
	ID          int
	Title       string
	Abstract    string
	Content     string
	ColumnID    int
	CategoryID  int
	TagIDs      []int
	CoverUrl    string
	Status      int
	IsTop       bool
	ViewCount   int
	LikeCount   int
	CreateTime  time.Time
	PublishTime time.Time
	UpdateTime  time.Time
	Version     int
}

func NewArticle(title, abstract, content, coverUrl string, columnID, categoryID int, tagIDs []int) (*Article, error) {
	if title == "" {
		return nil, fmt.Errorf("文章标题不能为空")
	}
	now := time.Now()
	return &Article{
		Title:      title,
		Abstract:   abstract,
		Content:    content,
		ColumnID:   columnID,
		CategoryID: categoryID,
		TagIDs:     tagIDs,
		CoverUrl:   coverUrl,
		CreateTime: now,
		UpdateTime: now,
	}, nil
}

func (a *Article) Publish() error {
	if a.Title == "" {
		return fmt.Errorf("文章标题不能为空")
	}

	if a.Abstract == "" {
		return fmt.Errorf("文章摘要不能为空")
	}

	if a.Content == "" {
		return fmt.Errorf("文章内容不能为空")
	}

	if a.CategoryID == 0 {
		return fmt.Errorf("文章分类不能为空")
	}

	a.Status = 1
	a.PublishTime = time.Now()
	a.UpdateTime = time.Now()
	a.Version++
	return nil
}

func (a *Article) Update(title, abstract, content string, columnID, categoryID int, tagIDs []int, coverUrl string) error {
	if a.ID == 0 {
		return fmt.Errorf("文章ID不能为空")
	}
	a.Title = title
	a.Abstract = abstract
	a.Content = content
	a.ColumnID = columnID
	a.CategoryID = categoryID
	a.TagIDs = tagIDs
	a.CoverUrl = coverUrl
	a.UpdateTime = time.Now()
	a.Version++
	return nil
}

func (a *Article) UpdateTopStatus(isTop bool) error {
	if a.Status != 1 {
		return fmt.Errorf("只有已发布的文章才能设置置顶状态")
	}
	a.IsTop = isTop
	a.Version++
	return nil
}

func (a *Article) IncrementViewCount() {
	a.ViewCount++
	a.Version++
}
