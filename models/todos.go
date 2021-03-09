package models

import (
	"github.com/99-66/go-gin-project-template/config"
	//_ "github.com/go-sql-driver/mysql"
)

func GetAllTodos() (todo []Todo, err error) {
	return todo, config.DB.Find(&todo).Error
}

func CreateTodo(todo *Todo) error {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(id int) (todo Todo, err error) {
	return todo, config.DB.First(&todo, id).Error
}

func UpdateTodo(todo, update_todo *Todo, id int) error {
	if err := config.DB.Model(&todo).Updates(update_todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *Todo, id int) error {
	if err := config.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}