package models

// Order represents a trade order
// @Description Trade order model
type Order struct {
    ID        uint    `gorm:"primaryKey"`
    Symbol    string  `json:"symbol"`
    Price     float64 `json:"price"`
    Quantity  int     `json:"quantity"`
    OrderType string  `json:"order_type"`
}
