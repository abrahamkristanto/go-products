package main

import (
	"github.com/abrahamkristanto/go-products/database"
	"github.com/abrahamkristanto/go-products/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/products", handlers.GetProducts)

	r.Run(":8080")
}
