package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/waldin-exe/golang-bootstrap/utils/response"
)

type middlewareManager struct {
	jwtSecret string
}

func NewMiddlewareManager(jwtSecret string) *middlewareManager {
	return &middlewareManager{
		jwtSecret: jwtSecret,
	}
}

func (m *middlewareManager) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.NewResponseUnauthorized(c, "missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == "" {
			return response.NewResponseUnauthorized(c, "invalid token format")
		}

		secret := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
			}
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			return response.NewResponseUnauthorized(c, "invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return response.NewResponseUnauthorized(c, "invalid token claims")
		}

		// Validasi exp otomatis seharusnya udah jalan di jwt.Parse
		// Tapi untuk jaga-jaga, bisa validasi manual juga (opsional)
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return response.NewResponseUnauthorized(c, "token expired")
			}
		}

		// Set ke context
		if email, ok := claims["email"].(string); ok {
			c.Locals("email", email)
		}
		if role, ok := claims["role"].(string); ok {
			c.Locals("role", role)
		}
		if userID, ok := claims["user_id"].(float64); ok {
			c.Locals("user_id", int(userID))
		}

		return c.Next()
	}
}

func (m *middlewareManager) RequireRoles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return response.NewResponseUnauthorized(c, "unauthorized")
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return response.NewResponseForbidden(c, "forbidden: access denied")
	}
}
