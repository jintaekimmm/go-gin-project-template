package controllers

import (
	"github.com/99-66/go-gin-project-template/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodos(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func GetTodo(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(404)})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	todo, err = models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item Not Found"})
		return
	}

	var updateBody models.Todo
	err = c.ShouldBindJSON(&updateBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateTodo(&todo, &updateBody, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var todo models.Todo
	todo, err = models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item Not Found"})
		return
	}

	err = models.DeleteTodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "")
}