package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"ginserver/database"
	"ginserver/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env", ".env.local", "Path to the .env file")
	flag.Parse()

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.InitDB("./database.db")

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.GetUser)

	router.Run(":3000")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
