package service

import (
	"gt-blog/backend/internal/mvc/dto"
	"gt-blog/backend/internal/mvc/repository"

	"github.com/gin-gonic/gin"
)

type VisitorService struct {
	visitorRepo repository.VisitorRepository
}

func NewVisitorService(visitorRepo repository.VisitorRepository) *VisitorService {
	return &VisitorService{
		visitorRepo: visitorRepo,
	}
}
func (v *VisitorService) GetAuthorInfo(c *gin.Context, id uint) (*dto.AuthorInfoResp, error) {
	author, err := v.visitorRepo.GetAuthorByID(c, id)
	if err != nil {
		return nil, err
	}
	socialAccounts := make([]dto.SocialAccount, 0)
	for _, socialAccount := range author.SocialAccounts {
		socialAccounts = append(socialAccounts, dto.SocialAccount{
			ID:       socialAccount.ID,
			Platform: socialAccount.Platform,
			Url:      socialAccount.Url,
		})
	}
	resp := &dto.AuthorInfoResp{
		ID:             author.ID,
		Nickname:       author.Nickname,
		AvatarUrl:      author.AvatarUrl,
		Bio:            author.Bio,
		Email:          author.Email,
		Wechat:         author.Wechat,
		AboutMe:        author.AboutMe,
		AboutBlog:      author.AboutBlog,
		SocialAccounts: socialAccounts,
	}
	return resp, nil
}
