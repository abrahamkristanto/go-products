package handlers

import (
	"fmt"
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
		return
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

func CreateProduct(c *gin.Context) {
	var input *models.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Quantity:    input.Quantity,
	}

	result := database.Get().Select("name", "description", "quantity").Create(&product)

	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusNotFound,
			"data":       nil,
			"message":    "no data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       product,
		"message":    "success",
	})

}

func DeleteProduct(c *gin.Context) {
	var product *models.Product
	id := c.Param("id")

	result := database.Get().Delete(&product, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusNotFound,
			"data":       nil,
			"message":    "no data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       id,
		"message":    "deleted successfully",
	})
}
