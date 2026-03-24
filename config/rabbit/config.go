package rabbit

import env "github.com/waldin-exe/golang-bootstrap/config/env"

type RabbitMQConfig struct {
	URL string
}

func Load() RabbitMQConfig {
	return RabbitMQConfig{
		URL: env.GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
	}
}
