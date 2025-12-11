package panels

import (
	"fmt"
	"strings"

	"dk/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateMarketplacePanel creates the marketplace panel with transaction and profit views
func CreateMarketplacePanel(state *models.AppState) tview.Primitive {
	// Left panel - Transactions
	state.TransactionsView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			state.App.Draw()
		})
	state.TransactionsView.SetBorder(true).
		SetTitle(" [magenta]Recent Logs[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Right top panel - Profit Stats
	state.ProfitView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.ProfitView.SetBorder(true).
		SetTitle(" [green]Profit Dashboard[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Right bottom panel - Taxes
	state.TaxesList = tview.NewList().
		ShowSecondaryText(true).
		SetHighlightFullLine(true)
	state.TaxesList.SetBorder(true).
		SetTitle(" [red]Taxes & Bills[white] ").
		SetBorderColor(tcell.ColorWhite)

	UpdateMarketplaceViews(state)

	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.ProfitView, 0, 1, false).
		AddItem(state.TaxesList, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(state.TransactionsView, 0, 1, false).
		AddItem(rightPanel, 0, 1, false)

	return mainLayout
}

// UpdateMarketplaceViews refreshes marketplace panel views
func UpdateMarketplaceViews(state *models.AppState) {
	updateTransactionsView(state)
	updateProfitView(state)
	updateTaxesView(state)
	UpdateStatusBar(state)
}

func updateTransactionsView(state *models.AppState) {
	var builder strings.Builder

	if len(state.Transactions) == 0 {
		builder.WriteString("[gray]No logs yet...[white]\n")
	} else {
		// Show last 30 transactions
		start := 0
		if len(state.Transactions) > 30 {
			start = len(state.Transactions) - 30
		}
		for i := len(state.Transactions) - 1; i >= start; i-- {
			t := state.Transactions[i]
			if t.IsTaxPayment {
				// Red for tax payments
				builder.WriteString(fmt.Sprintf("[gray][%s][white][red]-$%d[white] [red]PAID: %s[white]\n",
					t.Time.Format("15:04:05"), -t.Profit, t.Product))
			} else {
				// Green for sales
				builder.WriteString(fmt.Sprintf("[gray][%s][white][green]+$%d[white] [cyan]%s[white] bought [yellow]%s[white] - [green]%dx[white]\n",
					t.Time.Format("15:04:05"), t.Profit, t.BuyerName, t.Product, t.Amount))
			}
		}
	}

	state.TransactionsView.SetText(builder.String())
}

func updateProfitView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[green]STATISTICS[white]\n\n")

	builder.WriteString("Total Profit: [green]$" + fmt.Sprintf("%d", state.TotalProfit) + "[white]\n")
	builder.WriteString("Total Sales: [cyan]" + fmt.Sprintf("%d", len(state.Transactions)) + "[white] transactions\n")

	if len(state.Transactions) > 0 {
		avg := state.TotalProfit / len(state.Transactions)
		builder.WriteString("Average Sale: [yellow]$" + fmt.Sprintf("%d", avg) + "[white]\n")
	} else {
		builder.WriteString("Average Sale: [gray]$0[white]\n")
	}

	// Product breakdown
	builder.WriteString("\n[cyan]INVENTORY[white]\n\n")
	for _, prod := range state.Products {
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n", prod.Name))
		builder.WriteString(fmt.Sprintf("  Stock: [cyan]%d[white] | Price: [green]$%d[white]\n\n", prod.Stock, prod.Price))
	}

	state.ProfitView.SetText(builder.String())
}

func updateTaxesView(state *models.AppState) {
	state.TaxesList.Clear()

	if len(state.Taxes) == 0 {
		state.TaxesList.AddItem("[gray]No bills yet... Lucky you!", "", 0, nil)
	} else {
		for i, tax := range state.Taxes {
			mainText := fmt.Sprintf("[red]%s [-:-:-][white] - [red]$%d", tax.Name, tax.Amount)
			secondaryText := fmt.Sprintf("[gray]%s", tax.Description)

			// Capture index for closure
			idx := i
			state.TaxesList.AddItem(mainText, secondaryText, 0, func() {
				// Update selected tax index when item changes
				state.SelectedTax = idx
			})
		}

		// Set the current selection
		if state.SelectedTax >= 0 && state.SelectedTax < len(state.Taxes) {
			state.TaxesList.SetCurrentItem(state.SelectedTax)
		}
	}
}
