package command

type CreateCategoryCommand struct {
	Name        string
	Description string
}

type UpdateCategoryCommand struct {
	ID          int
	Name        string
	Description string
}

type DeleteCategoryCommand struct {
	ID int
}