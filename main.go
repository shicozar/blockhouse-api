package main

import (
    "blockhouse-api/config"
    "blockhouse-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize database connection
    config.ConnectDatabase()

    // Setup router and routes
    router := gin.Default()
    routes.SetupRoutes(router)

    // Start the server
    router.Run("0.0.0.0:8080")
}
