package dao

type ArticleTagDAO struct {
	ID         int `gorm:"column:id;primaryKey;autoIncrement"`
	ArticleID  int `gorm:"column:article_id;not null"`
	TagID      int `gorm:"column:tag_id;not null"`
	CreateTime int `gorm:"column:create_time;not null"`
}

func (a *ArticleTagDAO) TableName() string {
	return "article_tags"
}
