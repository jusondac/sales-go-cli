package panels
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



























































































































}	state.InfoView.SetText(builder.String())	builder.WriteString("[red]q[white] Quit\n")	builder.WriteString("[yellow]←/→[white] Switch panels\n")	builder.WriteString("[gray]Navigation:[white]\n")	builder.WriteString("\n\n[gray]═══════════════════[white]\n")	}		builder.WriteString(fmt.Sprintf("[white]Stock:[white] [cyan]%d units[white]\n", ing.Stock))		builder.WriteString(fmt.Sprintf("[white]Price:[white] [green]$%d[white]\n", ing.Price))		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n\n", ing.Name))		ing := state.Ingredients[state.SelectedIngredient]		builder.WriteString("[cyan]═══════════════════[white]\n\n")		builder.WriteString("[cyan] SELECTED ITEM[white]\n")		builder.WriteString("\n[cyan]═══════════════════[white]\n")	if state.SelectedIngredient < len(state.Ingredients) {	}		}			builder.WriteString(fmt.Sprintf("[yellow]  %s[white]\n\n", val))			builder.WriteString(fmt.Sprintf("[white]%-18s[white]\n", key))		if val, ok := data[key]; ok {	for _, key := range keys {	keys := []string{"💰 Money", "📦 Total Profit", "🏪 Products", "📊 Total Sales"}	builder.WriteString("[cyan]═══════════════════[white]\n\n")	builder.WriteString("[cyan]   USER STATUS[white]\n")	builder.WriteString("[cyan]═══════════════════[white]\n")	var builder strings.Builder	}		"📊 Total Sales":  fmt.Sprintf("%d transactions", len(state.Transactions)),		"🏪 Products":     fmt.Sprintf("%d types", len(state.Products)),		"📦 Total Profit": fmt.Sprintf("$%d", state.TotalProfit),		"💰 Money":        fmt.Sprintf("$%d", state.UserMoney),	data := map[string]string{func updateInfoView(state *models.AppState) {}	state.PreparationView.SetText(builder.String())	}		}			builder.WriteString(fmt.Sprintf("   [gray]Price: $%-4d Stock: %-4d[white]\n\n", prod.Price, prod.Stock))			builder.WriteString(fmt.Sprintf("   [white]%-15s[white]\n", prod.Name))		} else {			builder.WriteString("\n")			}				builder.WriteString(fmt.Sprintf("     [cyan]• %s x%d[white]\n", ingName, qty))			for ingName, qty := range prod.Ingredients {			builder.WriteString("   [white]Requires:[white]\n")			builder.WriteString(fmt.Sprintf("   [green]Price: $%-4d Stock: %-4d[white]\n", prod.Price, prod.Stock))			builder.WriteString(fmt.Sprintf("[black:green] ▶ %-15s [white]\n", prod.Name))		if i == state.SelectedProduct {	for i, prod := range state.Products {	builder.WriteString("[white]Keys: [green](n)[white] new product | [red](x)[white] delete\n\n")	builder.WriteString("[white]Keys: [green](e/d)[white] navigate | [green](p)[white] prepare\n")	builder.WriteString("[green]═══ PRODUCTS ═══[white]\n\n")	var builder strings.Builderfunc updatePreparationView(state *models.AppState) {}	state.IngredientsListView.SetText(builder.String())	}		}			builder.WriteString(fmt.Sprintf("   [gray]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))			builder.WriteString(fmt.Sprintf("   [white]%-12s[white]\n", ing.Name))		} else {			builder.WriteString(fmt.Sprintf("   [yellow]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))			builder.WriteString(fmt.Sprintf("[black:yellow] ▶ %-12s [white]\n", ing.Name))		if i == state.SelectedIngredient {	for i, ing := range state.Ingredients {	builder.WriteString("[white]Keys: [yellow](s/w)[white] navigate | [yellow](b)[white] buy\n\n")	builder.WriteString("[yellow]═══ INGREDIENTS ═══[white]\n\n")	var builder strings.Builderfunc updateIngredientsView(state *models.AppState) {}	updateInfoView(state)	updatePreparationView(state)	updateIngredientsView(state)func UpdateBusinessViews(state *models.AppState) {// UpdateBusinessViews refreshes all business panel views}	return mainLayout		AddItem(rightPanel, 0, 1, false)		AddItem(centerPanel, 0, 1, false).		AddItem(leftPanel, 0, 1, false).	mainLayout := tview.NewFlex().		AddItem(state.InfoView, 0, 1, false)	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).		AddItem(state.PreparationView, 0, 1, false)	centerPanel := tview.NewFlex().SetDirection(tview.FlexRow).		AddItem(state.IngredientsListView, 0, 1, false)	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).	// Layout	UpdateBusinessViews(state)		SetBorderColor(tcell.NewRGBColor(0, 255, 255))		SetTitle(" [cyan]Info Panel[white] ").	state.InfoView.SetBorder(true).		SetScrollable(true)		SetDynamicColors(true).	state.InfoView = tview.NewTextView().	// Right panel - Info		SetBorderColor(tcell.ColorGreen)		SetTitle(" [green]Prepare Products[white] ").	state.PreparationView.SetBorder(true).		SetScrollable(true)		SetDynamicColors(true).	state.PreparationView = tview.NewTextView().