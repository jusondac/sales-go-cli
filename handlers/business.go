package handlers

import (
	"fmt"

	"dk/models"
	"dk/panels"
	"dk/utils"

	"github.com/rivo/tview"
)

// BuyIngredient purchases the selected ingredient
func BuyIngredient(state *models.AppState) {
	if state.SelectedIngredient >= len(state.Ingredients) {
		return
	}

	ing := &state.Ingredients[state.SelectedIngredient]
	if state.UserMoney >= ing.Price {
		state.UserMoney -= ing.Price
		ing.Stock += 10
		state.App.QueueUpdateDraw(func() {
			panels.UpdateBusinessViews(state)
		})
	}
}

// PrepareProduct creates a product from ingredients
func PrepareProduct(state *models.AppState) {
	if state.SelectedProduct >= len(state.Products) {
		return
	}

	prod := &state.Products[state.SelectedProduct]

	// Check if we have enough ingredients
	canPrepare := true
	for ingName, qtyNeeded := range prod.Ingredients {
		found := false
		for i := range state.Ingredients {
			if state.Ingredients[i].Name == ingName {
				if state.Ingredients[i].Stock < qtyNeeded {
					canPrepare = false
				}
				found = true
				break
			}
		}
		if !found {
			canPrepare = false
		}
	}

	if canPrepare {
		// Deduct ingredients
		for ingName, qtyNeeded := range prod.Ingredients {
			for i := range state.Ingredients {
				if state.Ingredients[i].Name == ingName {
					state.Ingredients[i].Stock -= qtyNeeded
					break
				}
			}
		}
		prod.Stock++
		state.App.QueueUpdateDraw(func() {
			panels.UpdateBusinessViews(state)
		})
	}
}

// DeleteProduct removes the selected product
func DeleteProduct(state *models.AppState) {
	if state.SelectedProduct >= len(state.Products) || len(state.Products) <= 1 {
		return
	}

	state.Products = append(state.Products[:state.SelectedProduct], state.Products[state.SelectedProduct+1:]...)
	if state.SelectedProduct >= len(state.Products) {
		state.SelectedProduct = len(state.Products) - 1
	}
	state.App.QueueUpdateDraw(func() {
		panels.UpdateBusinessViews(state)
	})
}

// ShowNewProductForm displays a form to create a new product
func ShowNewProductForm(state *models.AppState) {
	var productForm *tview.Form

	productForm = tview.NewForm().
		AddInputField("Product Name", "", 20, nil, nil).
		AddInputField("Price", "", 10, nil, nil).
		AddButton("Create", func() {
			nameField := productForm.GetFormItem(0).(*tview.InputField)
			priceField := productForm.GetFormItem(1).(*tview.InputField)

			name := nameField.GetText()
			var price int
			fmt.Sscanf(priceField.GetText(), "%d", &price)

			if name != "" && price > 0 {
				newProduct := models.Product{
					Name:        name,
					Ingredients: map[string]int{},
					Price:       price,
					Stock:       0,
				}
				state.Products = append(state.Products, newProduct)
			}

			state.Pages.RemovePage("form")
			state.App.QueueUpdateDraw(func() {
				panels.UpdateBusinessViews(state)
			})
		}).
		AddButton("Cancel", func() {
			state.Pages.RemovePage("form")
		})

	productForm.SetBorder(true).
		SetTitle(" [yellow]New Product[white] ").
		SetTitleAlign(tview.AlignCenter)

	state.Pages.AddPage("form", utils.Modal(productForm, 50, 10), true, true)
	state.App.SetFocus(productForm)
}
