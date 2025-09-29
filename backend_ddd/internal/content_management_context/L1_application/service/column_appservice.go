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
	"math"
)

type ColumnAppService struct {
	columnRepo          repository.ColumnRepository
	columnDomainService *service.ColumnDomainService
}

func NewColumnAppService(
	columnRepo repository.ColumnRepository,
	columnDomainService *service.ColumnDomainService) *ColumnAppService {
	return &ColumnAppService{
		columnRepo:          columnRepo,
		columnDomainService: columnDomainService,
	}
}

func (s *ColumnAppService) CreateColumn(ctx context.Context, cmd command.CreateColumnCommand) (int, error) {
	newColumn := model.NewColumn(cmd.Name, cmd.Description, cmd.CoverUrl)
	id, err := s.columnRepo.Add(ctx, newColumn)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *ColumnAppService) UpdateColumn(ctx context.Context, cmd command.UpdateColumnCommand) error {
	column, err := s.columnRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if column == nil {
		return errors.New("column not found")
	}
	if err := column.Update(cmd.Name, cmd.Description, cmd.CoverUrl); err != nil {
		return err
	}

	return s.columnRepo.Update(ctx, column)
}

func (s *ColumnAppService) DeleteColumn(ctx context.Context, cmd command.DeleteColumnCommand) error {
	err := s.columnRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ColumnAppService) GetColumnList(ctx context.Context, cmd query.GetColumnListQuery) (*result.ColumnListResult, error) {
	columns, total, err := s.columnRepo.FindByPage(ctx, cmd.Page, cmd.PageSize)
	if err != nil {
		return nil, err
	}
	items := make([]*result.ColumnDetailResult, 0, len(columns))
	for _, column := range columns {
		items = append(items, &result.ColumnDetailResult{
			ID:           column.ID,
			Name:         column.Name,
			Description:  column.Description,
			CoverUrl:     column.CoverUrl,
			CreateTime:   column.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   column.UpdateTime.Format("2006-01-02 15:04:05"),
			ArticleCount: int(column.ArticleCount),
			ViewCount:    int(column.ViewCount),
			LikeCount:    int(column.LikeCount),
		})
	}
	return &result.ColumnListResult{
			Total:     int(total),
			Page:      cmd.Page,
			PageSize:  cmd.PageSize,
			PageCount: int(math.Ceil(float64(total) / float64(cmd.PageSize))),
			List:      items,
		},
		nil

}

func (s *ColumnAppService) GetColumnDetail(ctx context.Context, cmd query.GetColumnDetailQuery) (*result.ColumnDetailResult, error) {
	column, err := s.columnRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if column == nil {
		return nil, errors.New("column not found")
	}

	return &result.ColumnDetailResult{
			ID:           column.ID,
			Name:         column.Name,
			Description:  column.Description,
			CoverUrl:     column.CoverUrl,
			CreateTime:   column.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   column.UpdateTime.Format("2006-01-02 15:04:05"),
			ArticleCount: int(column.ArticleCount),
			ViewCount:    int(column.ViewCount),
			LikeCount:    int(column.LikeCount),
		},
		nil
}
