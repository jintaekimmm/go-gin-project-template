package main

import (
	"github.com/99-66/go-gin-project-template/controllers"
	"github.com/99-66/go-gin-project-template/repositories"
	"github.com/99-66/go-gin-project-template/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initTodoAPI(db *gorm.DB) controllers.TodoAPI {
	wire.Build(
		repositories.ProvideTodoRepository,
		services.ProvideTodoService,
		controllers.ProvideTodoAPI)

	return controllers.TodoAPI{}
}