package services

import (
	"github.com/99-66/go-gin-project-template/models"
	"github.com/99-66/go-gin-project-template/repositories"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
}

func ProvideTodoService(t repositories.TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

// FindAll
func (t *TodoService) FindAll() ([]models.Todo, error) {
	return t.TodoRepository.FindAll()
}

// FindById
func (t *TodoService) FindById(id uint) (models.Todo, error) {
	return t.TodoRepository.FindById(id)
}

// Create
func (t *TodoService) Create(todo *models.Todo) error {
	err := t.TodoRepository.Create(todo)
	return err
}

// Update
func (t *TodoService) Update(todo, updateTodo *models.Todo) error {
	err := t.TodoRepository.Update(todo, updateTodo)
	return err
}

// DeleteById
func (t *TodoService) DeleteById(todo *models.Todo, id uint) error {
	err := t.TodoRepository.DeleteById(todo, id)
	return err
}

// FindAllCollection
func (t *TodoService) FindAllCollection() ([]models.Location, error) {
	collection, err := t.TodoRepository.FindAllCollection()
	if err != nil {
		return nil, err
	}

	return collection, err
}