package repositories

import (
	"github.com/99-66/go-gin-project-template/models"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

)

type TodoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) TodoRepository {
	return TodoRepository{DB: DB}
}

// FindAll
func (t *TodoRepository) FindAll() (todo []models.Todo, err error) {
	return todo, t.DB.Find(&todo).Error
}

// FindById
func (t *TodoRepository) FindById(id uint) (todo models.Todo, err error) {
	return todo, t.DB.First(&todo, id).Error
}


// Create
func (t *TodoRepository) Create(todo *models.Todo) error {
	if err := t.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// Update
func (t *TodoRepository) Update(todo, updateTodo *models.Todo) error {
	if err := t.DB.Model(&todo).Updates(updateTodo).Error; err != nil {
		return err
	}
	return nil
}

// DeleteById
func (t *TodoRepository) DeleteById(todo *models.Todo, id uint) error {
	if err := t.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}