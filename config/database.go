package config

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host string `env:"DB_HOST"`
	Port int	`env:"DB_PORT"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

// InitDB Database env 값을 읽어서 로드한다.
func InitDB() (*gorm.DB, error) {
	dbConfig := DatabaseConfig{}
	// Env Parsing
	if err := env.Parse(&dbConfig); err != nil{
		return nil, errors.New("could not load database configuration")
	}

	// Make Database URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	// Make Database Connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate()

	return db, nil
}

