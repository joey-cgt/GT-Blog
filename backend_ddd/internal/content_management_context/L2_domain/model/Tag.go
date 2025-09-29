package model

import "time"

type Tag struct {
	ID           int
	Name         string
	Description  string
	ArticleCount int
	CreateTime   time.Time
	UpdateTime   time.Time
	Version      int
}

func NewTag(name, description string) *Tag {
	return &Tag{
		Name:         name,
		Description:  description,
		ArticleCount: 0,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Version:      0, 
	}
}

func (t *Tag) Update(name, description string) error {
	t.Name = name
	t.Description = description
	t.Version++
	return nil
}

func (t *Tag) IncrementArticleCount() error {
	t.ArticleCount++
	t.Version++
	return nil
}

func (t *Tag) DecrementArticleCount() error {
	t.ArticleCount--
	t.Version++
	return nil
}
