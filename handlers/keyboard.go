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
			go BuyIngredient(state)
		case 'p':
			go PrepareProduct(state)
		case 'n':
			go ShowNewProductForm(state)
		case 'x':
			go DeleteProduct(state)
		case 's':
			state.SelectedIngredient = (state.SelectedIngredient + 1) % len(state.Ingredients)
			state.App.QueueUpdateDraw(func() {
				panels.UpdateBusinessViews(state)
			})
		case 'w':
			state.SelectedIngredient = (state.SelectedIngredient - 1 + len(state.Ingredients)) % len(state.Ingredients)
			state.App.QueueUpdateDraw(func() {
				panels.UpdateBusinessViews(state)
			})
		case 'e':
			state.SelectedProduct = (state.SelectedProduct - 1 + len(state.Products)) % len(state.Products)
			state.App.QueueUpdateDraw(func() {
				panels.UpdateBusinessViews(state)
			})
		case 'd':
			state.SelectedProduct = (state.SelectedProduct + 1) % len(state.Products)
			state.App.QueueUpdateDraw(func() {
				panels.UpdateBusinessViews(state)
			})
		}
		return event
	})
}

// SetupGlobalKeyboard configures global keyboard handlers
func SetupGlobalKeyboard(pages *tview.Pages, state *models.AppState) {
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Global navigation
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

		// Business panel controls (only when on business page)
		if state.CurrentPage == 0 {
			switch event.Rune() {
			case 'b':
				go BuyIngredient(state)
				return nil
			case 'p':
				go PrepareProduct(state)
				return nil
			case 'n':
				go ShowNewProductForm(state)
				return nil
			case 'x':
				go DeleteProduct(state)
				return nil
			case 's':
				go func() {
					state.SelectedIngredient = (state.SelectedIngredient + 1) % len(state.Ingredients)
					state.App.QueueUpdateDraw(func() {
						panels.UpdateBusinessViews(state)
					})
				}()
				return nil
			case 'w':
				go func() {
					state.SelectedIngredient = (state.SelectedIngredient - 1 + len(state.Ingredients)) % len(state.Ingredients)
					state.App.QueueUpdateDraw(func() {
						panels.UpdateBusinessViews(state)
					})
				}()
				return nil
			case 'e':
				go func() {
					state.SelectedProduct = (state.SelectedProduct - 1 + len(state.Products)) % len(state.Products)
					state.App.QueueUpdateDraw(func() {
						panels.UpdateBusinessViews(state)
					})
				}()
				return nil
			case 'd':
				go func() {
					state.SelectedProduct = (state.SelectedProduct + 1) % len(state.Products)
					state.App.QueueUpdateDraw(func() {
						panels.UpdateBusinessViews(state)
					})
				}()
				return nil
			}
		}

		return event
	})
}
