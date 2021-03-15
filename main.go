package main

import (
	"github.com/99-66/go-gin-project-template/config"
	_ "github.com/99-66/go-gin-project-template/docs"
	"github.com/99-66/go-gin-project-template/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
// @query.collection.format multi

// @title Todo API
// @description Todo microservice API
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	// CORS allows all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	// Project routes
	todoAPI := initTodoAPI(db)
	routes.TodoRoute(r, todoAPI)

	// Swagger Settings
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}