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

	// Create status bar
	statusBar := panels.CreateStatusBar(state)

	// Create both panels
	businessPanel := panels.CreateBusinessPanel(state)
	marketplacePanel := panels.CreateMarketplacePanel(state)

	// Add pages
	state.Pages.AddPage("business", businessPanel, true, true)
	state.Pages.AddPage("marketplace", marketplacePanel, true, false)

	// Main layout with status bar on top
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(statusBar, 3, 0, false).
		AddItem(state.Pages, 0, 1, true)

	// Setup global keyboard handlers (includes business panel controls)
	handlers.SetupGlobalKeyboard(state.Pages, state)

	// Start background goroutine for marketplace transactions
	go handlers.SimulateMarketplace(state)

	// Run the application
	if err := state.App.SetRoot(mainLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
