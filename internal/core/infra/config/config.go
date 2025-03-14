package config

import (
	httpHandler "github.com/jhrick/go-template-api/internal/core/infra/http/handler"
	"github.com/jhrick/go-template-api/internal/core/infra/utils/env"
)

type Config struct {
	Server struct {
		Port    string
		Handler *httpHandler.Handler
	}
	DB *DB
}

type DB struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

func New() *Config {
	port := ":" + env.GetEnvOrDefault("APP_PORT", "8080")
	handler := httpHandler.New()

	return &Config{
		Server: struct {
			Port    string
			Handler *httpHandler.Handler
		}{Port: port, Handler: handler},
		DB: &DB{
			Name:     env.GetEnvOrError("DATABASE_NAME"),
			Host:     env.GetEnvOrDefault("DATABASE_HOST", "localhost"),
			Port:     env.GetEnvOrDefault("DATABASE_PORT", "5432"),
			User:     env.GetEnvOrError("DATABASE_USER"),
			Password: env.GetEnvOrError("DATABASE_PASSWORD"),
		},
	}
}
