package redis

import env "github.com/waldin-exe/golang-bootstrap/config/env"

type RedisConfig struct {
	Host string
	Port string
}

func Load() RedisConfig {
	return RedisConfig{
		Host: env.GetEnv("REDIS_HOST", "localhost"),
		Port: env.GetEnv("REDIS_PORT", "6379"),
	}
}
