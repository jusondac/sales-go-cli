package models

import "time"

// Ingredient represents a raw material for products
type Ingredient struct {
	Name  string
	Price int
	Stock int
}

// Product represents a sellable item made from ingredients
type Product struct {
	Name        string
	Ingredients map[string]int // ingredient name -> quantity needed
	Price       int
	Stock       int
}

// Transaction represents a marketplace sale or payment
type Transaction struct {
	BuyerName string
	Product   string
	Amount    int
	Profit    int
	Time      time.Time
	IsTaxPayment bool
}

// Tax represents a payment demand
type Tax struct {
	Name        string
	Description string
	Amount      int
	Time        time.Time
}
