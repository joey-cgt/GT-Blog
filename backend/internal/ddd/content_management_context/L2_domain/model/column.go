package model

import "time"

type Column struct {
	ID           int       // 专栏ID
	Name         string    // 专栏名称
	Description  string    // 专栏简介
	CoverUrl     string    // 专栏封面图URL
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	ArticleCount int64     // 专栏文章总数
	ViewCount    int64     // 专栏文章总浏览量
	LikeCount    int64     // 专栏文章总点赞数
	Version      int
}

func NewColumn(name, description, coverUrl string) *Column {
	return &Column{
		Name:         name,
		Description:  description,
		CoverUrl:     coverUrl,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		ArticleCount: 0,
		ViewCount:    0,
		LikeCount:    0,
		Version:      0,
	}
}

func (c *Column) Update(name, description, coverUrl string) error {
	c.Name = name
	c.Description = description
	c.CoverUrl = coverUrl
	c.Version++
	c.UpdateTime = time.Now()
	return nil
}
