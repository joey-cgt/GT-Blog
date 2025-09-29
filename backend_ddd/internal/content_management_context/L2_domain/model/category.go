package model

import "time"

type Category struct {
	ID           int
	Name         string
	Description  string
	ArticleCount int
	CreateTime   time.Time
	UpdateTime   time.Time
	Version      int
}

func NewCategory(name, description string) *Category {
	return &Category{
		Name:         name,
		Description:  description,
		ArticleCount: 0,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Version:      0,
	}
}

func (c *Category) Update(name, description string) error {
	c.Name = name
	c.Description = description
	return nil
}

func (c *Category) IncrementArticleCount() error {
	c.ArticleCount++
	return nil
}

func (c *Category) DecrementArticleCount() error {
	c.ArticleCount--
	return nil
}
