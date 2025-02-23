package handlers

import (
    "net/http"
    "blockhouse-api/models"
    "blockhouse-api/config"
    "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&order)
    c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Find(&orders)
    c.JSON(http.StatusOK, orders)
}
