package config

import "os"

const (
	SERVER_ADDRESS string = "SERVER_ADDRESS"
	ADMIN_ADDRESS  string = "ADMIN_ADDRESS"
	NATS_ADDRESS   string = "NATS_ADDRESS"
)

type Config struct {
	AdminAddress  string
	ServerAddress string
	NatsAddress   string
}

// NewFromEnv creates a Config from the environment variables with sensible fallbacks
func NewFromEnv() *Config {
	serverAddress := withDefaultValue(SERVER_ADDRESS, ":8080")
	adminAddress := withDefaultValue(ADMIN_ADDRESS, ":8081")
	natsAddress := withDefaultValue(NATS_ADDRESS, "nats://nats:4222")

	return &Config{adminAddress, serverAddress, natsAddress}
}

func withDefaultValue(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
