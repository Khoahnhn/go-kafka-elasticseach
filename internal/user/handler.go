package user

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all users"})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get user " + id})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create user"})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Remove user " + id})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Update user " + id})
}
