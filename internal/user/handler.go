package user

import (
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/user/request"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	users, total, err := GetUsersService(page, pageSize)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      users,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"totalPage": int(math.Ceil(float64(total) / float64(pageSize))),
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := GetUserByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var createUserRequest request.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createUser, err := CreateUserService(createUserRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": createUser})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")
	err := DeleteUserService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUserRequest request.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := UpdateUserService(id, updateUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func SearchUser(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	filters := map[string]string{
		"wildcard":      c.Query("wildcard"),
		"email":         c.Query("email"),
		"created_after": c.Query("created_after"),
	}

	users, err := SearchUserService(query, filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
