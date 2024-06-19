package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Env = initConfig()

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string

	PublicHost string
	Port       string

	JWTExpirationInSeconds string
	JWTSecret              string
}

func initConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "ecom"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),

		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),

		JWTExpirationInSeconds: getEnv("JWT_EXP", strconv.Itoa(3600*24*7)),
		JWTSecret:              getEnv("JWT_SECRET", "j@wL7(e"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
