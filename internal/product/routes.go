package product

import "github.com/gin-gonic/gin"

func RegisterProductRoutes(apiGroup *gin.RouterGroup) {
	productGroup := apiGroup.Group("/product")
	{
		productGroup.GET("", GetProducts)
		productGroup.GET("/:id", GetProductByID)
		productGroup.POST("/:id", CreateProduct)
		productGroup.DELETE("/:id", RemoveProduct)
		productGroup.PUT("/:id", UpdateProduct)
	}
}
