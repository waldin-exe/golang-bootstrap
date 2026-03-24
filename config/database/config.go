package database

import (
	"fmt"

	env "github.com/waldin-exe/golang-bootstrap/config/env"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() DatabaseConfig {
	return DatabaseConfig{
		Driver:   env.GetEnv("DB_DRIVER", "postgres"),
		Host:     env.GetEnv("DB_HOSTNAME", "localhost"),
		Port:     env.GetEnv("DB_PORT", "5432"),
		User:     env.GetEnv("DB_USERNAME", "postgres"),
		Password: env.GetEnv("DB_PASSWORD", ""),
		Name:     env.GetEnv("DB_NAME", "postgres"),
		SSLMode:  env.GetEnv("DB_SSLMODE", "disable"),
	}
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

func (d *DatabaseConfig) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.SSLMode,
	)
}
