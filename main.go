package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// create a server
	server := gin.Default() // Default() return a pointer

	routes.RegisterRoutes(server)

	// run a server
	server.Run(":8080") // localhost:8080
}
