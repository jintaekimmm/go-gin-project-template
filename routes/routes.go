package routes

import (
	"github.com/99-66/go-gin-project-template/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/todo", controllers.GetTodos)
			v1.POST("/todo", controllers.CreateTodo)
			v1.GET("/todo/:id", controllers.GetTodo)
			v1.PUT("/todo/:id", controllers.UpdateTodo)
			v1.DELETE("/todo/:id", controllers.DeleteTodo)
		}
	}

	return r
}