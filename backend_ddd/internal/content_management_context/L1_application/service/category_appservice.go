package service

import (
	"context"
	"errors"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/command"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/query"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/dto/result"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/model"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/repository"
	"gt-blog/backend_ddd/internal/content_management_context/L2_domain/service"
)

type CategoryAppService struct {
	categoryRepo          repository.CategoryRepository
	categoryDomainService *service.CategoryDomainService
}

func NewCategoryAppService(
	categoryRepo repository.CategoryRepository,
	categoryDomainService *service.CategoryDomainService) *CategoryAppService {
	return &CategoryAppService{
		categoryRepo:          categoryRepo,
		categoryDomainService: categoryDomainService,
	}
}

func (s *CategoryAppService) CreateCategory(ctx context.Context, cmd command.CreateCategoryCommand) (int, error) {
	newCateGory := model.NewCategory(cmd.Name, cmd.Description)
	id, err := s.categoryRepo.Add(ctx, newCateGory)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CategoryAppService) UpdateCategory(ctx context.Context, cmd command.UpdateCategoryCommand) error {
	category, err := s.categoryRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("category not found")
	}
	if err := category.Update(cmd.Name, cmd.Description); err != nil {
		return err
	}
	return s.categoryRepo.Update(ctx, category)
}

func (s *CategoryAppService) DeleteCategory(ctx context.Context, cmd command.DeleteCategoryCommand) error {
	err := s.categoryRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryAppService) GetCategoryList(ctx context.Context) (*result.CategoryListResult, error) {
	categories, total, err := s.categoryRepo.FindByPage(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*result.CategoryDetailResult, 0, len(categories))
	for _, category := range categories {
		items = append(items, &result.CategoryDetailResult{
			ID:           category.ID,
			Name:         category.Name,
			Description:  category.Description,
			CreateTime:   category.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   category.UpdateTime.Format("2006-01-02 15:04:05"),
			ArticleCount: category.ArticleCount,
		})
	}

	return &result.CategoryListResult{
		Total: total,
		List:  items,
	}, nil
}

func (s *CategoryAppService) GetCategoryDetail(ctx context.Context, qry query.GetCategoryDetailQuery) (*result.CategoryDetailResult, error) {
	category, err := s.categoryRepo.FindByID(ctx, qry.ID)

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return &result.CategoryDetailResult{
		ID:           category.ID,
		Name:         category.Name,
		Description:  category.Description,
		CreateTime:   category.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   category.UpdateTime.Format("2006-01-02 15:04:05"),
		ArticleCount: category.ArticleCount,
	}, nil

}
