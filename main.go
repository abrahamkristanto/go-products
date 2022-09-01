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
	r.GET("/products/:id", handlers.GetProduct)
	r.POST("/products", handlers.CreateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.Run(":8080")
}
