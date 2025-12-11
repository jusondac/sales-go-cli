package panels

import (
	"fmt"
	"strings"

	"dk/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateAnalyticsPanel creates the analytics panel with price tracking
func CreateAnalyticsPanel(state *models.AppState) tview.Primitive {
	// Left panel - Ingredient Prices
	state.IngredientPricesView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.IngredientPricesView.SetBorder(true).
		SetTitle(" [yellow]Ingredient Prices[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Center panel - Product Prices
	state.ProductPricesView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.ProductPricesView.SetBorder(true).
		SetTitle(" [green]Product Prices[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Right panel - Price History Info
	state.PriceHistoryView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.PriceHistoryView.SetBorder(true).
		SetTitle(" [cyan]Market Info[white] ").
		SetBorderColor(tcell.ColorWhite)

	UpdateAnalyticsViews(state)

	// Layout
	mainLayout := tview.NewFlex().
		AddItem(state.IngredientPricesView, 0, 1, false).
		AddItem(state.ProductPricesView, 0, 1, false).
		AddItem(state.PriceHistoryView, 0, 1, false)

	return mainLayout
}

// UpdateAnalyticsViews refreshes all analytics panel views
func UpdateAnalyticsViews(state *models.AppState) {
	updateIngredientPricesView(state)
	updateProductPricesView(state)
	updatePriceHistoryView(state)
	UpdateStatusBar(state)
}

func updateIngredientPricesView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[yellow]INGREDIENT MARKET PRICES[white]\n\n")

	for _, ing := range state.Ingredients {
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n", ing.Name))
		builder.WriteString(fmt.Sprintf("  Current: [green]$%d[white]\n", ing.Price))
		builder.WriteString(fmt.Sprintf("  Range: [cyan]$%d[white] - [cyan]$%d[white]\n", ing.Floor, ing.Ceil))
		builder.WriteString(fmt.Sprintf("  Volatility: [magenta]±$%d/tick[white]\n", ing.Step))
		builder.WriteString(fmt.Sprintf("  Stock: [white]%d[white]\n\n", ing.Stock))
	}

	state.IngredientPricesView.SetText(builder.String())
}

func updateProductPricesView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[green]PRODUCT MARKET PRICES[white]\n\n")

	for _, prod := range state.Products {
		builder.WriteString(fmt.Sprintf("[green]%s[white]\n", prod.Name))
		builder.WriteString(fmt.Sprintf("  Current: [green]$%d[white]\n", prod.Price))
		builder.WriteString(fmt.Sprintf("  Range: [cyan]$%d[white] - [cyan]$%d[white]\n", prod.Floor, prod.Ceil))
		builder.WriteString(fmt.Sprintf("  Volatility: [magenta]±$%d/tick[white]\n", prod.Step))
		builder.WriteString(fmt.Sprintf("  Stock: [white]%d[white]\n\n", prod.Stock))
	}

	state.ProductPricesView.SetText(builder.String())
}

func updatePriceHistoryView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[cyan]MARKET VOLATILITY SYSTEM[white]\n\n")

	builder.WriteString("[white]Prices update every tick with:[white]\n")
	builder.WriteString("  new = clamp(old + rand(−step, +step), floor, ceil)\n\n")

	builder.WriteString("[yellow]Key Concepts:[white]\n")
	builder.WriteString("  [magenta]Step[white] = Max change per tick\n")
	builder.WriteString("  [cyan]Floor[white] = Minimum price\n")
	builder.WriteString("  [cyan]Ceil[white] = Maximum price\n\n")

	builder.WriteString("[gray]Prices drift naturally over time\n")
	builder.WriteString("creating a dynamic market.\n\n")

	// Calculate average prices
	avgIngredientPrice := 0
	if len(state.Ingredients) > 0 {
		total := 0
		for _, ing := range state.Ingredients {
			total += ing.Price
		}
		avgIngredientPrice = total / len(state.Ingredients)
	}

	avgProductPrice := 0
	if len(state.Products) > 0 {
		total := 0
		for _, prod := range state.Products {
			total += prod.Price
		}
		avgProductPrice = total / len(state.Products)
	}

	builder.WriteString("[yellow]Market Averages:[white]\n")
	builder.WriteString(fmt.Sprintf("  Ingredients: [green]$%d[white]\n", avgIngredientPrice))
	builder.WriteString(fmt.Sprintf("  Products: [green]$%d[white]\n", avgProductPrice))

	state.PriceHistoryView.SetText(builder.String())
}
