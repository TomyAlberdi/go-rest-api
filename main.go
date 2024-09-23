package main

import (
	"github.com/joho/godotenv"
	"github.com/tomyalberdi/go-rest-api/config"
	"github.com/tomyalberdi/go-rest-api/models"
	"github.com/tomyalberdi/go-rest-api/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	models.MigrateUsers(config.DB)
	models.MigratePosts(config.DB)
	r := routes.SetupRouter()
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
