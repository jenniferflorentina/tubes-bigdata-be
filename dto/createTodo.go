package dto

type CreateTodoDTO struct {
	Title       string   `json:"title" validate:"empty=false"`
	Description string   `json:"description"`
	SubTodo     []string `json:"subTodo" validate:"empty=false"`
}
