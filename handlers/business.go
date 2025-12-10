package handlers
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






































































































}	state.Pages.AddPage("form", utils.Modal(productForm, 50, 10), true, true)		SetTitleAlign(tview.AlignCenter)		SetTitle(" [yellow]New Product[white] ").	productForm.SetBorder(true).		})			state.App.Draw()			state.Pages.RemovePage("form")		AddButton("Cancel", func() {		}).			state.App.Draw()			state.Pages.RemovePage("form")			}				panels.UpdateBusinessViews(state)				state.Products = append(state.Products, newProduct)				}					Stock:       0,					Price:       price,					Ingredients: map[string]int{},					Name:        name,				newProduct := models.Product{			if name != "" && price > 0 {			fmt.Sscanf(priceField.GetText(), "%d", &price)			var price int			name := nameField.GetText()			priceField := productForm.GetFormItem(1).(*tview.InputField)			nameField := productForm.GetFormItem(0).(*tview.InputField)		AddButton("Create", func() {		AddInputField("Price", "", 10, nil, nil).		AddInputField("Product Name", "", 20, nil, nil).	productForm = tview.NewForm().		var productForm *tview.Formfunc ShowNewProductForm(state *models.AppState) {// ShowNewProductForm displays a form to create a new product}	state.App.Draw()	panels.UpdateBusinessViews(state)	}		state.SelectedProduct = len(state.Products) - 1	if state.SelectedProduct >= len(state.Products) {	state.Products = append(state.Products[:state.SelectedProduct], state.Products[state.SelectedProduct+1:]...)	}		return	if state.SelectedProduct >= len(state.Products) || len(state.Products) <= 1 {func DeleteProduct(state *models.AppState) {// DeleteProduct removes the selected product}	}		state.App.Draw()		panels.UpdateBusinessViews(state)		prod.Stock++		}			}				}					break					state.Ingredients[i].Stock -= qtyNeeded				if state.Ingredients[i].Name == ingName {			for i := range state.Ingredients {		for ingName, qtyNeeded := range prod.Ingredients {		// Deduct ingredients	if canPrepare {	}		}			canPrepare = false		if !found {		}			}				break				found = true				}					canPrepare = false				if state.Ingredients[i].Stock < qtyNeeded {			if state.Ingredients[i].Name == ingName {		for i := range state.Ingredients {		found := false	for ingName, qtyNeeded := range prod.Ingredients {	canPrepare := true	// Check if we have enough ingredients	prod := &state.Products[state.SelectedProduct]	}		return	if state.SelectedProduct >= len(state.Products) {func PrepareProduct(state *models.AppState) {// PrepareProduct creates a product from ingredients}	}		state.App.Draw()		panels.UpdateBusinessViews(state)		ing.Stock += 10