package models
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

// Transaction represents a marketplace sale
type Transaction struct {






}	Time      time.Time	Profit    int	Amount    int	Product   string	BuyerName string