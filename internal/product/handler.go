package product

import "github.com/gin-gonic/gin"

func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all products"})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get product " + id})
}

func CreateProduct(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create product"})
}

func RemoveProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Remove product " + id})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Update product " + id})
}
