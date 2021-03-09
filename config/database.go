package config

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host string `env:"DB_HOST"`
	Port int	`env:"DB_PORT"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

// init Database env 값을 읽어서 로드한다.
// 정상적으로 환경변수 값을 로드했다면 connection url 을 생성하여 반환한다
func InitDB() (string, error) {
	dbConfig := DatabaseConfig{}
	if err := env.Parse(&dbConfig); err != nil{
		return "", errors.New("could not load database configuration")
	}

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	return dbUrl, nil
}

