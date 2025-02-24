package handlers

import (
    "net/http"
    "blockhouse-api/models"
    "blockhouse-api/config"
    "github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure for error responses
type ErrorResponse struct {
    Error string `json:"error"`
}

// CreateOrder creates a new trade order
// @Summary Create a new trade order
// @Description Submit a new trade order with symbol, price, quantity, and type
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order Details"
// @Success 201 {object} models.Order
// @Failure 400 {object} ErrorResponse
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&order)
    c.JSON(http.StatusCreated, order)
}

// GetOrders retrieves all orders
// @Summary Get all trade orders
// @Description Fetch a list of submitted trade orders
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func GetOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Find(&orders)
    c.JSON(http.StatusOK, orders)
}
