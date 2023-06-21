package main

import (
	"log"
	"os"

	"github.com/dsferreira54/golang-crud-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(os.Getenv("ROUTER_ADDR")); err != nil {
		log.Fatal(err)
	}
}
