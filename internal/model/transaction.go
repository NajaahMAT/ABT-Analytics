package model

import "time"

type Transaction struct {
	TransactionID   string    `json:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date"`
	UserID          string    `json:"user_id"`
	Country         string    `json:"country"`
	Region          string    `json:"region"`
	ProductID       string    `json:"product_id"`
	ProductName     string    `json:"product_name"`
	Category        string    `json:"category"`
	Price           float64   `json:"price"`
	Quantity        int       `json:"quantity"`
	TotalPrice      float64   `json:"total_price"`
	StockQuantity   int       `json:"stock_quantity"`
	AddedDate       time.Time `json:"added_date"`
}
