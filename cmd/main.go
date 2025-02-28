package main

import (
	"fmt"
	"github.com/Khoahnhn/go-kafka-elastichsearch/api"
	"github.com/Khoahnhn/go-kafka-elastichsearch/settings/env"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load biến môi trường
	env.LoadEnv()

	//tao router Gin
	router := gin.Default()

	// tao group API voi prefix /api/v1
	apiV1 := router.Group(env.GetEnv("APP_PREFIX", "api/v1"))

	api.RegisterRoutes(apiV1)

	// khoi dong server
	port := env.GetEnv("APP_PORT", "8080")
	fmt.Println("🚀 Server đang chạy tại port:", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Không thể khởi động server:", err)
	}
}
