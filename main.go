package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kishansuvarna09/go-url-shortner/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setup_routers(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}

func setup_routers(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenUrl)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.PUT("/api/v1/:shortID", routes.EditUrl)
	router.DELETE("/api/v1/:shortID", routes.DeleteUrl)
	router.POST("/api/v1/tag", routes.AddTag)
}
