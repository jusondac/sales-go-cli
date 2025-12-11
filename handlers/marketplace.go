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
	taxNames := []string{
		"Property Tax", "Business License", "Health Inspection", "Fire Safety Fee",
		"City Tax", "State Tax", "Rent Payment", "Utilities Bill",
		"Insurance Premium", "Waste Management", "Parking Permit", "Signage Fee",
	}
	taxDescriptions := map[string]string{
		"Property Tax":      "Annual property tax assessment",
		"Business License":  "Quarterly business operating license",
		"Health Inspection": "Monthly health and safety inspection",
		"Fire Safety Fee":   "Fire department safety compliance",
		"City Tax":          "Municipal business tax",
		"State Tax":         "State revenue tax",
		"Rent Payment":      "Commercial property rent",
		"Utilities Bill":    "Electricity, water, and gas",
		"Insurance Premium": "Business liability insurance",
		"Waste Management":  "Garbage and recycling service",
		"Parking Permit":    "Employee parking permits",
		"Signage Fee":       "Outdoor signage permit renewal",
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	lastTaxTime := time.Now()

	for range ticker.C {
		// Random tax generation (every 10-20 seconds)
		if time.Since(lastTaxTime).Seconds() > float64(10+rand.Intn(10)) {
			taxName := taxNames[rand.Intn(len(taxNames))]
			taxAmount := (50 + rand.Intn(450)) * 10 // $500 - $5000

			tax := models.Tax{
				Name:        taxName,
				Description: taxDescriptions[taxName],
				Amount:      taxAmount,
				Time:        time.Now(),
			}
			state.Taxes = append(state.Taxes, tax)
			lastTaxTime = time.Now()

			state.App.QueueUpdateDraw(func() {
				panels.UpdateMarketplaceViews(state)
			})
		}

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

// PayTax pays the selected tax bill
func PayTax(state *models.AppState) {
	if len(state.Taxes) == 0 {
		return
	}

	// Ensure SelectedTax is within bounds
	if state.SelectedTax >= len(state.Taxes) {
		state.SelectedTax = len(state.Taxes) - 1
	}
	if state.SelectedTax < 0 {
		state.SelectedTax = 0
	}

	tax := state.Taxes[state.SelectedTax]

	// Pay tax (can go into debt/negative)
	state.UserMoney -= tax.Amount

	// Add tax payment to transaction log
	transaction := models.Transaction{
		BuyerName:    "SYSTEM",
		Product:      tax.Name,
		Amount:       1,
		Profit:       -tax.Amount,
		Time:         time.Now(),
		IsTaxPayment: true,
	}
	state.Transactions = append(state.Transactions, transaction)

	state.Taxes = append(state.Taxes[:state.SelectedTax], state.Taxes[state.SelectedTax+1:]...)

	// Adjust selection after deletion
	if len(state.Taxes) == 0 {
		state.SelectedTax = 0
	} else if state.SelectedTax >= len(state.Taxes) {
		state.SelectedTax = len(state.Taxes) - 1
	}

	state.App.QueueUpdateDraw(func() {
		panels.UpdateMarketplaceViews(state)
		panels.UpdateBusinessViews(state)
	})
}
