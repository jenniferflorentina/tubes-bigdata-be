package dto

type CreateTodoDTO struct {
	Title       string   `json:"title" validate:"empty=false"`
	Description string   `json:"description"`
	SubTodo     []string `json:"subTodo" validate:"empty=false"`
	Deadline    string   `json:"deadline"`
	Status      bool     `json:"status"`
}
