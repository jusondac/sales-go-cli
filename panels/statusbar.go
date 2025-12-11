package panels

import (
	"fmt"

	"dk/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateStatusBar creates a single-line status bar spanning the full width
func CreateStatusBar(state *models.AppState) *tview.TextView {
	state.StatusBarView = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	
	state.StatusBarView.SetBorder(true).
		SetBorderColor(tcell.ColorWhite)

	UpdateStatusBar(state)
	return state.StatusBarView
}

// UpdateStatusBar refreshes the status bar content
func UpdateStatusBar(state *models.AppState) {
	totalSales := 0
	for _, t := range state.Transactions {
		if !t.IsTaxPayment {
			totalSales++
		}
	}

	text := fmt.Sprintf("[yellow]Money: $%d[white]  |  [green]Profit: $%d[white]  |  [cyan]Sales: %d[white]  |  [magenta]Products: %d[white]  |  [gray]Transactions: %d[white]",
		state.UserMoney, state.TotalProfit, totalSales, len(state.Products), len(state.Transactions))
	
	state.StatusBarView.SetText(text)
}

