package repository

import (
	"gt-blog/backend/internal/mvc/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VisitorRepository interface {
	GetAuthorByID(c *gin.Context, id uint) (*model.Admin, error)
}

type MySQLVisitorRepository struct {
	db *gorm.DB
}

func NewMySQLVisitorRepository(db *gorm.DB) *MySQLVisitorRepository {
	return &MySQLVisitorRepository{db: db}
}

func (r *MySQLVisitorRepository) GetAuthorByID(c *gin.Context, id uint) (*model.Admin, error) {
	var author model.Admin
	if err := r.db.WithContext(c).Preload("SocialAccounts").Where("id = ?", id).First(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}
