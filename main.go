package main

import (
	"github.com/99-66/go-gin-project-template/config"
	"github.com/99-66/go-gin-project-template/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//var err error

func main() {
	dsn, err := config.InitDB()
	if err != nil {
		log.Fatal("Database Initialize failed.", err)
	}

	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection failed.", err)
	}
	sqlDB, err := config.DB.DB()
	defer sqlDB.Close()

	//config.DB.AutoMigrate()

	r := routes.InitRouter()
	r.Run()
}