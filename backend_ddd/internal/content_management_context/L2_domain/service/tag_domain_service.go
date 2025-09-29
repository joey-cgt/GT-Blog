package service

import "gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"

type TagDomainService struct {
	tagRepository repository.TagRepository
}

func NewTagDomainService(tagRepository repository.TagRepository) *TagDomainService {
	return &TagDomainService{
		tagRepository: tagRepository,
	}
}
