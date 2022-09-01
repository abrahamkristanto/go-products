package handlers

import (
	"net/http"

	"github.com/abrahamkristanto/go-products/database"
	"github.com/abrahamkristanto/go-products/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []*models.Product

	result := database.Get().Find(&products)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusNotFound,
			"data":       nil,
			"message":    "no data",
		})
	}

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       products,
		"message":    "success",
	})
}
