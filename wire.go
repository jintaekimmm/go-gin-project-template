package main

import (
	"github.com/99-66/go-gin-project-template/controllers"
	"github.com/99-66/go-gin-project-template/repositories"
	"github.com/99-66/go-gin-project-template/services"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func initTodoAPI(db *gorm.DB, mongodb *mongo.Client) controllers.TodoAPI {
	wire.Build(
		repositories.ProvideTodoRepository,
		services.ProvideTodoService,
		controllers.ProvideTodoAPI)

	return controllers.TodoAPI{}
}