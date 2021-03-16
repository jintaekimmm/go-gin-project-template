package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type MongoDBConfig struct {
	Host string `env:"MONGODB_HOST"`
	Port int	`env:"MONGODB_PORT"`
	Name string `env:"MONGODB_NAME"`
	User string `env:"MONGODB_USER"`
	Password string `env:"MONGODB_PASSWORD"`
	SSL bool `env:"MONGODB_SSL"`
}

// InitDB Database Connection을 생성하여 로드한다.
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

// InitMongoDB MongoDB Connection을 생성하여 반환한다
func InitMongoDB() (*mongo.Client, error) {
	mongoConfig := MongoDBConfig{}
	// Env Parsing
	if err := env.Parse(&mongoConfig); err != nil {
		return nil, errors.New("cloud not load mongodb configuration")
	}
	// Make MongoDB URL
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/?ssl=%t",
		mongoConfig.User,
		mongoConfig.Password,
		mongoConfig.Host,
		mongoConfig.Port,
		mongoConfig.SSL)

	// Make MongoDB Connection
	clientOption := options.Client().ApplyURI(dsn)
	db, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}

	// Check Connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}