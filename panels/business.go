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
			builder.WriteString(fmt.Sprintf("[yellow]▶ %s:$%d [%d][white]\n", ing.Name, ing.Price, ing.Stock))
		} else {
			builder.WriteString(fmt.Sprintf("  %s:$%d [%d]\n", ing.Name, ing.Price, ing.Stock))
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
			builder.WriteString(fmt.Sprintf("[green]▶ %s - $%d [%d][white]\n", prod.Name, prod.Price, prod.Stock))
		} else {
			builder.WriteString(fmt.Sprintf("  %s - $%d [%d]\n", prod.Name, prod.Price, prod.Stock))
		}
	}

	state.PreparationView.SetText(builder.String())
}

func updateInfoView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[cyan]USER STATUS[white]\n")
	builder.WriteString(fmt.Sprintf("Money: [yellow]$%d[white]\n", state.UserMoney))
	builder.WriteString(fmt.Sprintf("Profit: [green]$%d[white]\n", state.TotalProfit))
	builder.WriteString(fmt.Sprintf("Products: [white]%d types[white]\n", len(state.Products)))
	builder.WriteString(fmt.Sprintf("Sales: [white]%d transactions[white]\n", len(state.Transactions)))

	if state.SelectedProduct < len(state.Products) {
		builder.WriteString("\n[cyan]PRODUCT INFO[white]\n")
		prod := state.Products[state.SelectedProduct]
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n", prod.Name))
		builder.WriteString(fmt.Sprintf("Price: [green]$%d[white]\n", prod.Price))
		builder.WriteString(fmt.Sprintf("Stock: [cyan]%d[white]\n", prod.Stock))
		if len(prod.Ingredients) > 0 {
			builder.WriteString("Requires:\n")
			for ingName, qty := range prod.Ingredients {
				builder.WriteString(fmt.Sprintf("  %s x%d\n", ingName, qty))
			}
		}
	}

	builder.WriteString("\n[gray]Navigation:[white]\n")
	builder.WriteString("[yellow]←/→[white] Switch panels\n")
	builder.WriteString("[red]q[white] Quit\n")

	state.InfoView.SetText(builder.String())
}
