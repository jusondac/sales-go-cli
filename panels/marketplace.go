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
		SetTitle(" [magenta]Marketplace Transactions[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Right panel - Profit Stats
	state.ProfitView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.ProfitView.SetBorder(true).
		SetTitle(" [green]Profit Dashboard[white] ").
		SetBorderColor(tcell.ColorWhite)

	UpdateMarketplaceViews(state)

	mainLayout := tview.NewFlex().
		AddItem(state.TransactionsView, 0, 2, false).
		AddItem(state.ProfitView, 0, 1, false)

	return mainLayout
}

// UpdateMarketplaceViews refreshes marketplace panel views
func UpdateMarketplaceViews(state *models.AppState) {
	updateTransactionsView(state)
	updateProfitView(state)
}

func updateTransactionsView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[magenta]═══ RECENT SALES ═══[white]\n\n")

	if len(state.Transactions) == 0 {
		builder.WriteString("[gray]No transactions yet...[white]\n")
	} else {
		// Show last 20 transactions
		start := 0
		if len(state.Transactions) > 20 {
			start = len(state.Transactions) - 20
		}
		for i := len(state.Transactions) - 1; i >= start; i-- {
			t := state.Transactions[i]
			builder.WriteString(fmt.Sprintf("[gray][%s][white][green]+$%d[white] [cyan]%s[white] bought [yellow]%s[white] - [green]%dx[white]\n",
				t.Time.Format("15:04:05"), t.Profit, t.BuyerName, t.Product, t.Amount))
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
