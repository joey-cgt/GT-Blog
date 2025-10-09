package command

type CreateTagCommand struct {
	Name        string
	Description string
}

type UpdateTagCommand struct {
	ID          int
	Name        string
	Description string
}

type DeleteTagCommand struct {
	ID int
}