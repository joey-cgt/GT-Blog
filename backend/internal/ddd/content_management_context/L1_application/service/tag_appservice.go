package service

import (
	"context"
	"errors"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/command"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/query"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/dto/result"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/model"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/repository"
	"gt-blog/backend/internal/ddd/content_management_context/L2_domain/service"
)

type TagAppService struct {
	tagRepo          repository.TagRepository
	tagDomainService *service.TagDomainService
}

func NewTagAppService(
	tagRepo repository.TagRepository,
	tagDomainService *service.TagDomainService) *TagAppService {
	return &TagAppService{
		tagRepo:          tagRepo,
		tagDomainService: tagDomainService,
	}
}

func (s *TagAppService) CreateTag(ctx context.Context, cmd command.CreateTagCommand) (int, error) {
	newTag := model.NewTag(cmd.Name, cmd.Description)
	id, err := s.tagRepo.Add(ctx, newTag)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TagAppService) UpdateTag(ctx context.Context, cmd command.UpdateTagCommand) error {
	tag, err := s.tagRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if tag == nil {
		return errors.New("tag not found")
	}
	if err := tag.Update(cmd.Name, cmd.Description); err != nil {
		return err
	}
	return s.tagRepo.Update(ctx, tag)
}

func (s *TagAppService) DeleteTag(ctx context.Context, cmd command.DeleteTagCommand) error {
	err := s.tagRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *TagAppService) GetTagList(ctx context.Context, qry query.GetTagListQuery) (*result.TagListResult, error) {
	tags, total, err := s.tagRepo.FindByPage(ctx, qry.Page, qry.PageSize)
	if err != nil {
		return nil, err
	}
	items := make([]*result.TagDetailResult, 0, len(tags))
	for _, tag := range tags {
		items = append(items, &result.TagDetailResult{
			ID:           tag.ID,
			Name:         tag.Name,
			Description:  tag.Description,
			CreateTime:   tag.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   tag.UpdateTime.Format("2006-01-02 15:04:05"),
			ArticleCount: int(tag.ArticleCount),
		})
	}

	return &result.TagListResult{
		Items: items,
		Total: int(total),
	}, nil

}

func (s *TagAppService) GetTagDetail(ctx context.Context, cmd query.GetTagDetailQuery) (*result.TagDetailResult, error) {
	panic("...")
}
