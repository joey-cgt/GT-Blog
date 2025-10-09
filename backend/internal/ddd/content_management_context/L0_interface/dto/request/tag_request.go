package request

type CreateTagRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
}

type UpdateTagRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=1,max=20"`
	Description string `json:"description" validate:"required,min=1,max=100"`
}

type DeleteTagRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetTagDetailRequest struct {
	ID int `uri:"id" validate:"required"`
}

type GetTagListRequest struct {
	Page     int `form:"page" validate:"required,min=1"`
	PageSize int `form:"pageSize" validate:"required,min=1,max=100"`
}
