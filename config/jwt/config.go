package jwt

import env "github.com/waldin-exe/golang-bootstrap/config/env"

type JWTConfig struct {
	AccessSecret        string
	RefreshSecret       string
	AccessExpireMinutes int
	RefreshExpireDays   int
}

func Load() JWTConfig {
	return JWTConfig{
		AccessSecret:        env.GetEnv("JWT_SECRET_KEY", "your-access-secret"),
		RefreshSecret:       env.GetEnv("JWT_REFRESH_SECRET_KEY", "your-refresh-secret"),
		AccessExpireMinutes: env.GetEnvInt("JWT_ACCESS_TOKEN_EXPIRE_MINUTES", 15),
		RefreshExpireDays:   env.GetEnvInt("JWT_REFRESH_TOKEN_EXPIRE_DAYS", 7),
	}
}
