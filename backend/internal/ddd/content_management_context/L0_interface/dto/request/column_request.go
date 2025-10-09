package request

type CreateColumnRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
	CoverUrl    string `json:"coverUrl" validate:"required,min=1,max=100"`
}

type UpdateColumnRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
	CoverUrl    string `json:"coverUrl" validate:"required,min=1,max=100"`
}

type DeleteColumnRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetColumnDetailRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetColumnListRequest struct {
	Page     int `form:"page" validate:"required,min=1"`
	PageSize int `form:"pageSize" validate:"required,min=1,max=100"`
}
