package models

type Todo struct {
	Id uint `json:"id"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *Todo) TableName() string {
	return "todo"
}

type Location struct {
	ID string `bson:"_id" json:"_id"`
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}