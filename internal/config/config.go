package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Variable struct {
	Port       string
	DBUrl      string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	RedisURL   string
	JWTSecret  string
}

func LoadEnv() *Variable {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load enviroment variables")
	}

	return &Variable{
		Port:       os.Getenv("PORT"),
		DBUrl:      os.Getenv("DATABASE_URL"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		RedisURL:   os.Getenv("REDIS_URL"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
