package api

import (
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/product"
	"github.com/Khoahnhn/go-kafka-elastichsearch/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	// routes health
	router.GET("/health", HealthCheck)

	user.RegisterUserRoutes(router)
	product.RegisterProductRoutes(router)
}
