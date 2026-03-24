package app

import env "github.com/waldin-exe/golang-bootstrap/config/env"

type AppConfig struct {
	Name      string
	Port      string
	Env       string
	JWTSecret string
}

func Load() AppConfig {
	return AppConfig{
		Name:      env.GetEnv("APP_NAME", "go-bootstrap"),
		Port:      env.GetEnv("APP_PORT", "2000"),
		Env:       env.GetEnv("APP_ENV", "development"),
		JWTSecret: env.GetEnv("JWT_SECRET_KEY", "secret"),
	}
}
