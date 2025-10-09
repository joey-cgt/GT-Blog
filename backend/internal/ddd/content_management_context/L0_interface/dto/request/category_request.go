package request

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
	CoverUrl    string `json:"coverUrl" validate:"required,min=1,max=255"`
}

type UpdateCategoryRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
	CoverUrl    string `json:"coverUrl" validate:"required,min=1,max=255"`
}

type DeleteCategoryRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetCategoryDetailRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetCategoryListRequest struct {
	Page     int `form:"page" validate:"required,min=1"`
	PageSize int `form:"pageSize" validate:"required,min=1,max=100"`
}
