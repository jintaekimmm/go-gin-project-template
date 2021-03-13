package main

import (
	"github.com/99-66/go-gin-project-template/config"
	"github.com/99-66/go-gin-project-template/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := initRoutes(db)
	log.Fatal(r.Run())
}


func initRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	todoAPI := initTodoAPI(db)
	routes.TodoRoute(r, todoAPI)

	return r
}