package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/pkg/db"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/routes"
)

func main() {
	r := gin.Default()

	db := db.GetDatabase()

	routes.InitRoutes(r, db)

	fmt.Println("🚀 Server running at http://localhost:8080 🚀")
	r.Run(":8080")
}
