package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(apiGroup *gin.RouterGroup) {
	userGroup := apiGroup.Group("/user")
	{
		userGroup.GET("", GetUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", CreateUser)
		userGroup.DELETE("/:id", RemoveUser)
		userGroup.PUT("/:id", UpdateUser)
	}
}
