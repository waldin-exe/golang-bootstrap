package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/waldin-exe/golang-bootstrap/config/app"
	"github.com/waldin-exe/golang-bootstrap/config/database"
	"github.com/waldin-exe/golang-bootstrap/config/jwt"
	"github.com/waldin-exe/golang-bootstrap/config/rabbit"
	"github.com/waldin-exe/golang-bootstrap/config/redis"
)

type Config struct {
	App      app.AppConfig
	Database database.DatabaseConfig
	Redis    redis.RedisConfig
	RabbitMQ rabbit.RabbitMQConfig
	JWT      jwt.JWTConfig
}

func LoadConfig() *Config {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("[Config] No .env file found, using system env")
	}

	return &Config{
		App:      app.Load(),
		Database: database.Load(),
		Redis:    redis.Load(),
		RabbitMQ: rabbit.Load(),
		JWT:      jwt.Load(),
	}
}
