package main

import (
	"math/rand"
	"time"

	"dk/handlers"
	"dk/models"
	"dk/panels"

	"github.com/rivo/tview"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize application state
	state := models.NewAppState()
	state.App = tview.NewApplication()
	state.Pages = tview.NewPages()

	// Create both panels
	businessPanel := panels.CreateBusinessPanel(state)
	marketplacePanel := panels.CreateMarketplacePanel(state)

	// Add pages
	state.Pages.AddPage("business", businessPanel, true, true)
	state.Pages.AddPage("marketplace", marketplacePanel, true, false)

	// Setup global keyboard handlers (includes business panel controls)
	handlers.SetupGlobalKeyboard(state.Pages, state)

	// Start background goroutine for marketplace transactions
	go handlers.SimulateMarketplace(state)

	// Run the application
	if err := state.App.SetRoot(state.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
