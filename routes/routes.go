package routes

import (
    "github.com/gin-gonic/gin"
    "blockhouse-api/handlers"
	"github.com/swaggo/gin-swagger"
     swaggerFiles "github.com/swaggo/files"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/orders", handlers.CreateOrder)
    router.GET("/orders", handlers.GetOrders)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}