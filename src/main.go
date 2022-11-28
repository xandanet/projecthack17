package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"podcast/src/app"
	"podcast/src/zlog"
)

type Config struct {
	DB *sql.DB
}

// main godoc
// @title BEIS AAR API
// @version 1.0
// @description Proof of Concept for BEIS AAR API
// @host localhost:5000
// @BasePath /
func main() {
	zlog.InitializeLogger()

	path, _ := os.Getwd()

	envPath := os.Getenv("ENV_PATH")
	if envPath == "" {
		envPath = fmt.Sprintf("%s/.env", path)
	}

	err := godotenv.Load(envPath)
	if err != nil && os.Getenv("APP_ENV") != "local" {
		zlog.Logger.Fatal(fmt.Sprintf("error loading .env file: %s", err))
	}

	appCred := app.Credentials{
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBName:  os.Getenv("DB_NAME"),
		DBName2: os.Getenv("DB_NAME_2"),
		DBTls:   os.Getenv("DB_TLS"),
		AppPort: os.Getenv("APP_PORT"),
	}

	app.StartApplication(&appCred)
}
