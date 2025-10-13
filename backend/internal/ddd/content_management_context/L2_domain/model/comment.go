package model

import "time"

type Comment struct {
	ID        uint       `json:"id"`
	Nickname  string     `json:"nickname"`
	Email     string     `json:"email"`
	Content   string     `json:"content"`
	ArticleID uint       `json:"article_id"`
	ParentID  *uint      `json:"parent_id"`
	Children  []*Comment `json:"children"`
	CreatedAt time.Time  `json:"created_at"`
}
