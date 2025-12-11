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
		SetBorderColor(tcell.ColorWhite)

	// Center panel - Prepare Products
	state.PreparationView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.PreparationView.SetBorder(true).
		SetTitle(" [green]Prepare Products[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Product Info panel
	state.ProductInfoView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.ProductInfoView.SetBorder(true).
		SetTitle(" [magenta]Product Info[white] ").
		SetBorderColor(tcell.ColorWhite)

	// Right panel - Info
	state.InfoView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	state.InfoView.SetBorder(true).
		SetTitle(" [cyan]Info Panel[white] ").
		SetBorderColor(tcell.ColorWhite)

	UpdateBusinessViews(state)

	// Layout
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.IngredientsListView, 0, 1, false)

	centerPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(state.PreparationView, 0, 1, false).
		AddItem(state.ProductInfoView, 0, 1, false)

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
	updateProductInfoView(state)
	updateInfoView(state)
	UpdateStatusBar(state)
}

func updateIngredientsView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[white]Keys: [yellow](s/w)[white] navigate | [yellow](b)[white] buy\n\n")

	for i, ing := range state.Ingredients {
		if i == state.SelectedIngredient {
			builder.WriteString(fmt.Sprintf("[black:yellow]%s:$%d [%d][-:-:-]\n", ing.Name, ing.Price, ing.Stock))
		} else {
			builder.WriteString(fmt.Sprintf("[white:-:-]%s:$%d [%d]\n", ing.Name, ing.Price, ing.Stock))
		}
	}

	state.IngredientsListView.SetText(builder.String())
}

func updatePreparationView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[white]Keys: [green](e/d)[white] navigate | [green](p)[white] prepare\n")
	builder.WriteString("[white]Keys: [green](n)[white] new product | [red](x)[white] delete\n\n")

	for i, prod := range state.Products {
		if i == state.SelectedProduct {
			builder.WriteString(fmt.Sprintf("[black:green]%s - $%d [%d][-:-:-]\n", prod.Name, prod.Price, prod.Stock))
		} else {
			builder.WriteString(fmt.Sprintf("[white:-:-]%s - $%d [%d]\n", prod.Name, prod.Price, prod.Stock))
		}
	}

	state.PreparationView.SetText(builder.String())
}

func updateProductInfoView(state *models.AppState) {
	var builder strings.Builder

	if state.SelectedProduct < len(state.Products) {
		prod := state.Products[state.SelectedProduct]
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n\n", prod.Name))
		builder.WriteString(fmt.Sprintf("Price: [green]$%d[white]\n", prod.Price))
		builder.WriteString(fmt.Sprintf("Stock: [cyan]%d[white]\n", prod.Stock))
		if len(prod.Ingredients) > 0 {
			builder.WriteString("\nRequires:\n")
			for ingName, qty := range prod.Ingredients {
				builder.WriteString(fmt.Sprintf("  %s x%d\n", ingName, qty))
			}
		}
	} else {
		builder.WriteString("[gray]No product selected[white]")
	}

	state.ProductInfoView.SetText(builder.String())
}

func updateInfoView(state *models.AppState) {
	var builder strings.Builder
	builder.WriteString("[cyan]USER STATUS[white]\n")
	builder.WriteString(fmt.Sprintf("Money: [yellow]$%d[white]\n", state.UserMoney))
	builder.WriteString(fmt.Sprintf("Profit: [green]$%d[white]\n", state.TotalProfit))
	builder.WriteString(fmt.Sprintf("Products: [white]%d types[white]\n", len(state.Products)))
	builder.WriteString(fmt.Sprintf("Sales: [white]%d transactions[white]\n", len(state.Transactions)))

	builder.WriteString("\n[gray]Navigation:[white]\n")
	builder.WriteString("[yellow]←/→[white] Switch panels\n")
	builder.WriteString("[red]q[white] Quit\n")

	state.InfoView.SetText(builder.String())
}
