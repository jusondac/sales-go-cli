package handlers
package handlers

import (
	"dk/models"
	"dk/panels"

	"github.com/gdamore/tcell/v2"
)

// SetupBusinessKeyboard configures keyboard handlers for the business panel
func SetupBusinessKeyboard(layout *tview.Flex, state *models.AppState) {
	layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'b':
			BuyIngredient(state)
		case 'p':
			PrepareProduct(state)
		case 'n':
			ShowNewProductForm(state)
		case 'x':
			DeleteProduct(state)
		case 's':
			state.SelectedIngredient = (state.SelectedIngredient + 1) % len(state.Ingredients)
			panels.UpdateBusinessViews(state)
		case 'w':
			state.SelectedIngredient = (state.SelectedIngredient - 1 + len(state.Ingredients)) % len(state.Ingredients)
			panels.UpdateBusinessViews(state)
		case 'e':
			state.SelectedProduct = (state.SelectedProduct - 1 + len(state.Products)) % len(state.Products)
			panels.UpdateBusinessViews(state)
		case 'd':

























}	})		return event		}			return nil			state.App.Stop()		} else if event.Rune() == 'q' {			return nil			pages.SwitchToPage("marketplace")			state.CurrentPage = 1		} else if event.Key() == tcell.KeyRight {			return nil			pages.SwitchToPage("business")			state.CurrentPage = 0		if event.Key() == tcell.KeyLeft {	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {func SetupGlobalKeyboard(pages *tview.Pages, state *models.AppState) {// SetupGlobalKeyboard configures global keyboard handlers}	})		return event		}			panels.UpdateBusinessViews(state)			state.SelectedProduct = (state.SelectedProduct + 1) % len(state.Products)