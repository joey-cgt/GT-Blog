package service

import "gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"

type CategoryDomainService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryDomainService(categoryRepository repository.CategoryRepository) *CategoryDomainService {
	return &CategoryDomainService{
		categoryRepository: categoryRepository,
	}
}
