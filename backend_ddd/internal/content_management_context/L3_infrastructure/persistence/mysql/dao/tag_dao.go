package dao

import "time"

type TagDAO struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text"`
	CreateTime  time.Time `gorm:"column:create_time;not null"`
	UpdateTime  time.Time `gorm:"column:update_time;not null"`
	Version     int       `gorm:"column:version;not null;default:0"`
}

func (t *TagDAO) TableName() string {
	return "tags"
}
