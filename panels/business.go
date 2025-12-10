package panels

import (
	"fmt"
	"strings"

	"dk/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateBusinessPanel creates the business panel with three sub-panels
func CreateBusinessPanel(state *models.AppState) tview.Primitive {
	// Left panel - Buy Ingredients
	state.IngredientsListView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.IngredientsListView.SetBorder(true).
		SetTitle(" [yellow]Buy Ingredients[white] ").
		SetBorderColor(tcell.ColorYellow)

	// Center panel - Prepare Products
	state.PreparationView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.PreparationView.SetBorder(true).
		SetTitle(" [green]Prepare Products[white] ").
		SetBorderColor(tcell.ColorGreen)

	// Right panel - Info
	state.InfoView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.InfoView.SetBorder(true).
		SetTitle(" [cyan]Info Panel[white] ").
		SetBorderColor(tcell.NewRGBColor(0, 255, 255))

	UpdateBusinessViews(state)

	// Layout
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.IngredientsListView, 0, 1, false)

	centerPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.PreparationView, 0, 1, false)

	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.InfoView, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(leftPanel, 0, 1, false).
		AddItem(centerPanel, 0, 1, false).
		AddItem(rightPanel, 0, 1, false)

	return mainLayout
}

// UpdateBusinessViews refreshes all business panel views
func UpdateBusinessViews(state *models.AppState) {
	updateIngredientsView(state)
	updatePreparationView(state)
	updateInfoView(state)
}

func updateIngredientsView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[yellow]═══ INGREDIENTS ═══[white]\n\n")
	builder.WriteString("[white]Keys: [yellow](s/w)[white] navigate | [yellow](b)[white] buy\n\n")

	for i, ing := range state.Ingredients {
		if i == state.SelectedIngredient {
			builder.WriteString(fmt.Sprintf("[yellow]▶ %-12s[white]\n", ing.Name))
			builder.WriteString(fmt.Sprintf("  [yellow]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))
		} else {
			builder.WriteString(fmt.Sprintf("  %-12s\n", ing.Name))
			builder.WriteString(fmt.Sprintf("  [gray]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))
		}
	}

	state.IngredientsListView.SetText(builder.String())
}

func updatePreparationView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[green]═══ PRODUCTS ═══[white]\n\n")
	builder.WriteString("[white]Keys: [green](e/d)[white] navigate | [green](p)[white] prepare\n")
	builder.WriteString("[white]Keys: [green](n)[white] new product | [red](x)[white] delete\n\n")

	for i, prod := range state.Products {
		if i == state.SelectedProduct {
			builder.WriteString(fmt.Sprintf("[green]▶ %-15s[white]\n", prod.Name))
			builder.WriteString(fmt.Sprintf("  [green]Price: $%-4d Stock: %-4d[white]\n", prod.Price, prod.Stock))
			builder.WriteString("  Requires:\n")
			for ingName, qty := range prod.Ingredients {
				builder.WriteString(fmt.Sprintf("    [cyan]• %s x%d[white]\n", ingName, qty))
			}
			builder.WriteString("\n")
		} else {
			builder.WriteString(fmt.Sprintf("  %-15s\n", prod.Name))
			builder.WriteString(fmt.Sprintf("  [gray]Price: $%-4d Stock: %-4d[white]\n\n", prod.Price, prod.Stock))
		}
	}

	state.PreparationView.SetText(builder.String())
}

func updateInfoView(state *models.AppState) {
	data := map[string]string{
		"💰 Money":        fmt.Sprintf("$%d", state.UserMoney),
		"📦 Total Profit": fmt.Sprintf("$%d", state.TotalProfit),
		"🏪 Products":     fmt.Sprintf("%d types", len(state.Products)),
		"📊 Total Sales":  fmt.Sprintf("%d transactions", len(state.Transactions)),
	}

	var builder strings.Builder
	builder.WriteString("[cyan]═══════════════════[white]\n")
	builder.WriteString("[cyan]   USER STATUS[white]\n")
	builder.WriteString("[cyan]═══════════════════[white]\n\n")

	keys := []string{"💰 Money", "📦 Total Profit", "🏪 Products", "📊 Total Sales"}
	for _, key := range keys {
		if val, ok := data[key]; ok {
			builder.WriteString(fmt.Sprintf("[white]%-18s[white]\n", key))
			builder.WriteString(fmt.Sprintf("[yellow]  %s[white]\n\n", val))
		}
	}

	if state.SelectedIngredient < len(state.Ingredients) {
		builder.WriteString("\n[cyan]═══════════════════[white]\n")
		builder.WriteString("[cyan] SELECTED ITEM[white]\n")
		builder.WriteString("[cyan]═══════════════════[white]\n\n")

		ing := state.Ingredients[state.SelectedIngredient]
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n\n", ing.Name))
		builder.WriteString(fmt.Sprintf("[white]Price:[white] [green]$%d[white]\n", ing.Price))
		builder.WriteString(fmt.Sprintf("[white]Stock:[white] [cyan]%d units[white]\n", ing.Stock))
	}

	builder.WriteString("\n\n[gray]═══════════════════[white]\n")
	builder.WriteString("[gray]Navigation:[white]\n")
	builder.WriteString("[yellow]←/→[white] Switch panels\n")
	builder.WriteString("[red]q[white] Quit\n")

	state.InfoView.SetText(builder.String())
}
