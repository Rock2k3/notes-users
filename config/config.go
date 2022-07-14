package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	httpAddress   string
	grpcPort      string
	datasourceUrl string
}

func NewAppConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &AppConfig{
		httpAddress:   fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		grpcPort:      fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")),
		datasourceUrl: os.Getenv("APP_DATASOURCE_URL"),
	}, nil
}

func (c *AppConfig) HttpAddress() string {
	return c.httpAddress
}

func (c *AppConfig) GrpcPort() string {
	return c.grpcPort
}

func (c *AppConfig) DatasourceUrl() string {
	return c.datasourceUrl
}
