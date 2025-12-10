package handlers

import (
	"math/rand"
	"time"

	"dk/models"
	"dk/panels"
	"dk/utils"
)

// SimulateMarketplace runs a background goroutine that simulates customer purchases
func SimulateMarketplace(state *models.AppState) {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Find products with stock
		availableProducts := []int{}
		for i, prod := range state.Products {
			if prod.Stock > 0 {
				availableProducts = append(availableProducts, i)
			}
		}

		if len(availableProducts) == 0 {
			continue
		}

		// Random purchase
		prodIdx := availableProducts[rand.Intn(len(availableProducts))]
		prod := &state.Products[prodIdx]
		amount := rand.Intn(utils.Min(3, prod.Stock)) + 1

		if prod.Stock >= amount {
			prod.Stock -= amount
			profit := prod.Price * amount
			state.TotalProfit += profit
			state.UserMoney += profit

			transaction := models.Transaction{
				BuyerName: names[rand.Intn(len(names))],
				Product:   prod.Name,
				Amount:    amount,
				Profit:    profit,
				Time:      time.Now(),
			}
			state.Transactions = append(state.Transactions, transaction)

			state.App.QueueUpdateDraw(func() {
				panels.UpdateMarketplaceViews(state)
				if state.CurrentPage == 0 {
					panels.UpdateBusinessViews(state)
				}
			})
		}
	}
}
