package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Order struct
type Order struct {
    ID        uint    `gorm:"primaryKey"`
    Symbol    string  `json:"symbol"`
    Price     float64 `json:"price"`
    Quantity  int     `json:"quantity"`
    OrderType string  `json:"order_type"`
}

var DB *gorm.DB

// Connects to SQLite
func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&Order{})
}

// Create Order
func CreateOrder(c *gin.Context) {
    var order Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    DB.Create(&order)
    c.JSON(http.StatusCreated, order)
}

// Get Orders
func GetOrders(c *gin.Context) {
    var orders []Order
    DB.Find(&orders)
    c.JSON(http.StatusOK, orders)
}

func main() {
    ConnectDatabase()

    router := gin.Default()
    router.POST("/orders", CreateOrder)
    router.GET("/orders", GetOrders)

    router.Run(":8080")
}
