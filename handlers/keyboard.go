package handlers

import (
	"dk/models"
	"dk/panels"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
			state.SelectedProduct = (state.SelectedProduct + 1) % len(state.Products)
			panels.UpdateBusinessViews(state)
		}
		return event
	})
}

// SetupGlobalKeyboard configures global keyboard handlers
func SetupGlobalKeyboard(pages *tview.Pages, state *models.AppState) {
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyLeft {
			state.CurrentPage = 0
			pages.SwitchToPage("business")
			return nil
		} else if event.Key() == tcell.KeyRight {
			state.CurrentPage = 1
			pages.SwitchToPage("marketplace")
			return nil
		} else if event.Rune() == 'q' {
			state.App.Stop()
			return nil
		}
		return event
	})
}
