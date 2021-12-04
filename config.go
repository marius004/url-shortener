package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	JwtSecret  string
	DbPassword string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbName     = os.Getenv("DB_NAME")
		jwtSecret  = os.Getenv("JWT_SECRET")
		dbPassword = os.Getenv("DB_PASSWORD")
	)

	return &Config{
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbName:     dbName,
		JwtSecret:  jwtSecret,
		DbPassword: dbPassword,
	}
}
