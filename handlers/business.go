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
		// Mark as initialized for volatility system
		ing.IsInitialized = true
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
		// Mark as initialized for volatility system
		prod.IsInitialized = true
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
	var productName string
	var productPrice int

	// Step 1: Get product name and price
	productForm = tview.NewForm().
		AddInputField("Product Name", "", 20, nil, nil).
		AddInputField("Price", "", 10, nil, nil).
		AddButton("Next", func() {
			go func() {
				nameField := productForm.GetFormItem(0).(*tview.InputField)
				priceField := productForm.GetFormItem(1).(*tview.InputField)

				productName = nameField.GetText()
				fmt.Sscanf(priceField.GetText(), "%d", &productPrice)

				if productName == "" || productPrice <= 0 {
					return
				}

				// Step 2: Select ingredients
				state.App.QueueUpdateDraw(func() {
					state.Pages.RemovePage("form")
					showIngredientSelectionForm(state, productName, productPrice)
				})
			}()
		}).
		AddButton("Cancel", func() {
			go func() {
				state.App.QueueUpdateDraw(func() {
					state.Pages.RemovePage("form")
				})
			}()
		})

	productForm.SetBorder(true).
		SetTitle(" [yellow]New Product - Step 1[white] ").
		SetTitleAlign(tview.AlignCenter)

	state.Pages.AddPage("form", utils.Modal(productForm, 50, 10), true, true)
	state.App.SetFocus(productForm)
}

func showIngredientSelectionForm(state *models.AppState, productName string, productPrice int) {
	var ingredientForm *tview.Form
	selectedIngredients := make(map[string]int)

	ingredientForm = tview.NewForm()

	// Add input fields for each ingredient
	for _, ing := range state.Ingredients {
		ingredientForm.AddInputField(ing.Name, "0", 10, nil, nil)
	}

	ingredientForm.AddButton("Create", func() {
		go func() {
			// Collect ingredient quantities
			for i, ing := range state.Ingredients {
				inputField := ingredientForm.GetFormItem(i).(*tview.InputField)
				var qty int
				fmt.Sscanf(inputField.GetText(), "%d", &qty)
				if qty > 0 {
					selectedIngredients[ing.Name] = qty
				}
			}

			newProduct := models.Product{
				Name:          productName,
				Ingredients:   selectedIngredients,
				Price:         productPrice,
				Stock:         0,
				Step:          10,
				Floor:         productPrice * 7 / 10,
				Ceil:          productPrice * 15 / 10,
				IsInitialized: false,
			}
			state.Products = append(state.Products, newProduct)

			state.App.QueueUpdateDraw(func() {
				state.Pages.RemovePage("form")
				panels.UpdateBusinessViews(state)
			})
		}()
	}).
		AddButton("Cancel", func() {
			go func() {
				state.App.QueueUpdateDraw(func() {
					state.Pages.RemovePage("form")
				})
			}()
		})

	ingredientForm.SetBorder(true).
		SetTitle(" [yellow]New Product - Step 2: Select Ingredients[white] ").
		SetTitleAlign(tview.AlignCenter)

	state.Pages.AddPage("form", utils.Modal(ingredientForm, 60, 15), true, true)
	state.App.SetFocus(ingredientForm)
}
