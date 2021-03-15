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

// FindAll godoc
// @Summary FindAll Todo List
// @Description FindAll Todo List
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {array} models.Todo
// @Failure 500 {object} config.APIError
// @Router /todo [get]
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

// Create godoc
// @Summary Create Todo
// @Description Create Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param todo body models.Todo true "Create Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} config.APIError
// @Failure 500 {object} config.APIError
// @Router /todo [post]
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

// FindById godoc
// @Summary FindById Todo
// @Description FindById Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} config.APIError
// @Failure 404 {object} config.APIError
// @Router /todo/{id} [get]
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
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update godoc
// @Summary Update Todo
// @Description Update Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Param todo body models.Todo true "Update Todo"
// @Success 200 {object} models.Todo
// @Failure 400 {object} config.APIError
// @Failure 404 {object} config.APIError
// @Failure 500 {object} config.APIError
// @Router /todo/{id} [put]
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

// Delete godoc
// @Summary Delete Todo
// @Description Delete Todo
// @Tags Todo
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id path integer true "ID"
// @Success 204 {string} {}
// @Failure 400 {object} config.APIError
// @Failure 404 {object} config.APIError
// @Failure 500 {object} config.APIError
// @Router /todo/{id} [delete]
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