package controllers

import (
	"github.com/99-66/go-gin-project-template/models"
	"github.com/99-66/go-gin-project-template/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoAPI struct {
	TodoService services.TodoService
}

func ProvideTodoAPI(t services.TodoService) TodoAPI {
	return TodoAPI{TodoService: t}
}

// FindAll
func (t *TodoAPI) FindAll(c *gin.Context) {
	// Usage Roles
	//roles, ok := c.Get("roles")
	//if !ok {
	//	return
	//}

	todos, err := t.TodoService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// Create
func (t *TodoAPI) Create(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = t.TodoService.Create(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// FindById
func (t *TodoAPI) FindById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(404)})
	}

	c.JSON(http.StatusOK, todo)
}

// Update
func (t *TodoAPI) Update(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(404)})
		return
	}

	var updateTodo models.Todo
	err = c.BindJSON(&updateTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = t.TodoService.Update(&todo, &updateTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Delete
func (t *TodoAPI) Delete(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.TodoService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(404)})
		return
	}

	err = t.TodoService.DeleteById(&todo, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}