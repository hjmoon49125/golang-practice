package main

import (
	"fmt"
	"os"
	"server_temp/source/src/database"
	"server_temp/source/src/router"

	"github.com/joho/godotenv"
)

// @title Swagger
// @version 1.0
// @description Swagger
// @host localhost:8080
// @securityDefinitions.apikey jwt
// @in header
// @name Authorization
// schemes http
func main() {
	initEnv()
	initDataBase()
	initServer()
}

func initEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func initDataBase() {
	if err := database.Init(); err != nil {
		panic(fmt.Errorf("database init failed"))
	}
}

func initServer() {
	router := router.NewRoute()
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(`:%s`, port))
}
