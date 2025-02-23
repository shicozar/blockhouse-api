package routes

import (
    "github.com/gin-gonic/gin"
    "blockhouse-api/handlers"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/orders", handlers.CreateOrder)
    router.GET("/orders", handlers.GetOrders)
}