package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetDataSourceURL() string {
	user := GetEnv("DB_USER", "root")
	password := GetEnv("DB_PASSWORD", "root")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "microservices")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}

func GetApplicationPort() int {
	portStr := GetEnv("APPLICATION_PORT", "8082")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 8082
	}
	return port
}