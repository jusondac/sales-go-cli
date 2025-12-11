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
	state.TaxesView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.TaxesView.SetBorder(true).
		SetTitle(" [red]Taxes & Bills[white] ").
		SetBorderColor(tcell.ColorWhite)

	UpdateMarketplaceViews(state)

	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.ProfitView, 0, 1, false).
		AddItem(state.TaxesView, 0, 1, false)

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
	var builder strings.Builder

	if len(state.Taxes) == 0 {
		builder.WriteString("[gray]No bills yet... Lucky you![white]\n")
	} else {
		builder.WriteString("[white]Keys: [yellow]↑/↓[white] navigate | [yellow](y)[white] pay\n\n")
		for i, tax := range state.Taxes {
			if i == state.SelectedTax {
				// White text on red background for selected tax
				builder.WriteString(fmt.Sprintf("[white:red] [%s] [%s] [$%d] [-:-:-]\n", tax.Name, tax.Description, tax.Amount))
			} else {
				// Red text for name and amount, gray for description
				builder.WriteString(fmt.Sprintf("[red:-:-] [%s] [white:-:-][gray:-:-][%s][white:-:-] [red:-:-][$%d][-:-:-]\n", tax.Name, tax.Description, tax.Amount))
			}
		}
	}

	state.TaxesView.SetText(builder.String())
}
