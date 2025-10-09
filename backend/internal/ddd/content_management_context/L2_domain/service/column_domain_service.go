package service

import "gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"

type ColumnDomainService struct {
	columnRepository repository.ColumnRepository
}

func NewColumnDomainService(columnRepository repository.ColumnRepository) *ColumnDomainService {
	return &ColumnDomainService{
		columnRepository: columnRepository,
	}
}
