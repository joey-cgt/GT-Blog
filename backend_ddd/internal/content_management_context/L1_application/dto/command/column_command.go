package command

type CreateColumnCommand struct {
	Name        string
	Description string
	CoverUrl    string
}

type UpdateColumnCommand struct {
	ID          int
	Name        string
	Description string
	CoverUrl    string
}

type DeleteColumnCommand struct {
	ID int
}
