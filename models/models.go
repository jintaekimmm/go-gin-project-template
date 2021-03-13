package models

type Todo struct {

	Id uint `json:"id"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *Todo) TableName() string {
	return "todo"
}