package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv 加载环境变量
func LoadEnv() error {
	return godotenv.Load()
}

// GetPort 获取端口号
func GetPort() string {
	return os.Getenv("PORT")
}

// GetDatabaseURL 获取数据库URL
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

// GetEnvironment 获取环境
func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	return env
}

// IsProduction 判断是否为生产环境
func IsProduction() bool {
	return GetEnvironment() == "production"
} 