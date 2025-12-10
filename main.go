package main

import (
	"fmt"
	"math/rand"
	"strings"
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

	// Create both panels
	businessPanel := panels.CreateBusinessPanel(state)
	marketplacePanel := panels.CreateMarketplacePanel(state)

	// Setup keyboard handlers for business panel
	if flex, ok := businessPanel.(*tview.Flex); ok {
		handlers.SetupBusinessKeyboard(flex, state)
	}

	// Add pages
	state.Pages.AddPage("business", businessPanel, true, true)
	state.Pages.AddPage("marketplace", marketplacePanel, true, false)

	// Setup global keyboard handlers
	handlers.SetupGlobalKeyboard(state.Pages, state)

	// Start background goroutine for marketplace transactions
	go handlers.SimulateMarketplace(state)

	// Run the application
	if err := state.App.SetRoot(state.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func initializeData() {
	ingredients = []Ingredient{
		{Name: "Flour", Price: 10, Stock: 0},
		{Name: "Sugar", Price: 15, Stock: 0},
		{Name: "Eggs", Price: 20, Stock: 0},
		{Name: "Butter", Price: 25, Stock: 0},
		{Name: "Milk", Price: 12, Stock: 0},
		{Name: "Chocolate", Price: 30, Stock: 0},
	}

	products = []Product{
		{
			Name: "Cake",
			Ingredients: map[string]int{
				"Flour":  3,
				"Sugar":  2,
				"Eggs":   2,
				"Butter": 1,
			},
			Price: 100,
			Stock: 0,
		},
		{
			Name: "Cookie",
			Ingredients: map[string]int{
				"Flour":  2,
				"Sugar":  1,
				"Butter": 1,
			},
			Price: 50,
			Stock: 0,
		},
	}

	transactions = []Transaction{}
}

func createBusinessPanel() tview.Primitive {
	// Left panel - Buy Ingredients
	ingredientsListView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	ingredientsListView.SetBorder(true).
		SetTitle(" [yellow]Buy Ingredients[white] ").
		SetBorderColor(tcell.ColorYellow)

	// Center panel - Prepare Products
	preparationView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	preparationView.SetBorder(true).
		SetTitle(" [green]Prepare Products[white] ").
		SetBorderColor(tcell.ColorGreen)

	// Right panel - Info
	infoView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	infoView.SetBorder(true).
		SetTitle(" [cyan]Info Panel[white] ").
		SetBorderColor(tcell.ColorCyan)

	updateBusinessViews()

	// Layout
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ingredientsListView, 0, 1, false)

	centerPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(preparationView, 0, 1, false)

	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoView, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(leftPanel, 0, 1, false).
		AddItem(centerPanel, 0, 1, false).
		AddItem(rightPanel, 0, 1, false)

	// Key handler for business operations
	mainLayout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'b':
			buyIngredient()
		case 'p':
			prepareProduct()
		case 'n':
			showNewProductForm()
		case 'd':
			deleteProduct()
		case 's':
			selectedIngredient = (selectedIngredient + 1) % len(ingredients)
			updateBusinessViews()
		case 'w':
			selectedIngredient = (selectedIngredient - 1 + len(ingredients)) % len(ingredients)
			updateBusinessViews()
		case 'e':
			selectedProduct = (selectedProduct - 1 + len(products)) % len(products)
			updateBusinessViews()
		case 'd':
			selectedProduct = (selectedProduct + 1) % len(products)
			updateBusinessViews()
		}
		return event
	})

	return mainLayout
}

func createMarketplacePanel() tview.Primitive {
	// Left panel - Transactions
	transactionsView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	transactionsView.SetBorder(true).
		SetTitle(" [magenta]Marketplace Transactions[white] ").
		SetBorderColor(tcell.ColorMagenta)

	// Right panel - Profit Stats
	profitView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true)
	profitView.SetBorder(true).
		SetTitle(" [green]Profit Dashboard[white] ").
		SetBorderColor(tcell.ColorGreen)

	updateMarketplaceViews()

	mainLayout := tview.NewFlex().
		AddItem(transactionsView, 0, 2, false).
		AddItem(profitView, 0, 1, false)

	return mainLayout
}

func updateBusinessViews() {
	// Update ingredients list
	var ingredientsBuilder strings.Builder
	ingredientsBuilder.WriteString("[yellow]═══ INGREDIENTS ═══[white]\n\n")
	ingredientsBuilder.WriteString("[white]Keys: [yellow](s/w)[white] navigate | [yellow](b)[white] buy\n\n")

	for i, ing := range ingredients {
		if i == selectedIngredient {
			ingredientsBuilder.WriteString(fmt.Sprintf("[black:yellow] ▶ %-12s [white]\n", ing.Name))
			ingredientsBuilder.WriteString(fmt.Sprintf("   [yellow]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))
		} else {
			ingredientsBuilder.WriteString(fmt.Sprintf("   [white]%-12s[white]\n", ing.Name))
			ingredientsBuilder.WriteString(fmt.Sprintf("   [gray]Price: $%-4d Stock: %-4d[white]\n\n", ing.Price, ing.Stock))
		}
	}

	ingredientsListView.SetText(ingredientsBuilder.String())

	// Update preparation view
	var prepBuilder strings.Builder
	prepBuilder.WriteString("[green]═══ PRODUCTS ═══[white]\n\n")
	prepBuilder.WriteString("[white]Keys: [green](e/d)[white] navigate | [green](p)[white] prepare\n")
	prepBuilder.WriteString("[white]Keys: [green](n)[white] new product | [red](d)[white] delete\n\n")

	for i, prod := range products {
		if i == selectedProduct {
			prepBuilder.WriteString(fmt.Sprintf("[black:green] ▶ %-15s [white]\n", prod.Name))
			prepBuilder.WriteString(fmt.Sprintf("   [green]Price: $%-4d Stock: %-4d[white]\n", prod.Price, prod.Stock))
			prepBuilder.WriteString("   [white]Requires:[white]\n")
			for ingName, qty := range prod.Ingredients {
				prepBuilder.WriteString(fmt.Sprintf("     [cyan]• %s x%d[white]\n", ingName, qty))
			}
			prepBuilder.WriteString("\n")
		} else {
			prepBuilder.WriteString(fmt.Sprintf("   [white]%-15s[white]\n", prod.Name))
			prepBuilder.WriteString(fmt.Sprintf("   [gray]Price: $%-4d Stock: %-4d[white]\n\n", prod.Price, prod.Stock))
		}
	}

	preparationView.SetText(prepBuilder.String())

	// Update info view
	updateDetailPanel(map[string]string{
		"💰 Money":        fmt.Sprintf("$%d", userMoney),
		"📦 Total Profit": fmt.Sprintf("$%d", totalProfit),
		"🏪 Products":     fmt.Sprintf("%d types", len(products)),
		"📊 Total Sales":  fmt.Sprintf("%d transactions", len(transactions)),
	})
}

func updateDetailPanel(data map[string]string) {
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

	if selectedIngredient < len(ingredients) {
		builder.WriteString("\n[cyan]═══════════════════[white]\n")
		builder.WriteString("[cyan] SELECTED ITEM[white]\n")
		builder.WriteString("[cyan]═══════════════════[white]\n\n")

		ing := ingredients[selectedIngredient]
		builder.WriteString(fmt.Sprintf("[yellow]%s[white]\n\n", ing.Name))
		builder.WriteString(fmt.Sprintf("[white]Price:[white] [green]$%d[white]\n", ing.Price))
		builder.WriteString(fmt.Sprintf("[white]Stock:[white] [cyan]%d units[white]\n", ing.Stock))
	}

	builder.WriteString("\n\n[gray]═══════════════════[white]\n")
	builder.WriteString("[gray]Navigation:[white]\n")
	builder.WriteString("[yellow]←/→[white] Switch panels\n")
	builder.WriteString("[red]q[white] Quit\n")

	infoView.SetText(builder.String())
}

func updateMarketplaceViews() {
	// Update transactions list
	var transBuilder strings.Builder
	transBuilder.WriteString("[magenta]═══ RECENT SALES ═══[white]\n\n")

	if len(transactions) == 0 {
		transBuilder.WriteString("[gray]No transactions yet...[white]\n")
	} else {
		// Show last 20 transactions
		start := 0
		if len(transactions) > 20 {
			start = len(transactions) - 20
		}
		for i := len(transactions) - 1; i >= start; i-- {
			t := transactions[i]
			transBuilder.WriteString(fmt.Sprintf("[cyan]%s[white] bought [yellow]%s[white] - [green]%dx[white]\n",
				t.BuyerName, t.Product, t.Amount))
			transBuilder.WriteString(fmt.Sprintf("  [green]+$%d[white] profit [gray]%s[white]\n\n",
				t.Profit, t.Time.Format("15:04:05")))
		}
	}

	transactionsView.SetText(transBuilder.String())

	// Update profit view
	var profitBuilder strings.Builder
	profitBuilder.WriteString("[green]═══ STATISTICS ═══[white]\n\n")

	profitBuilder.WriteString("[white]Total Profit:[white]\n")
	profitBuilder.WriteString(fmt.Sprintf("[green]  $%d[white]\n\n", totalProfit))

	profitBuilder.WriteString("[white]Total Sales:[white]\n")
	profitBuilder.WriteString(fmt.Sprintf("[cyan]  %d transactions[white]\n\n", len(transactions)))

	profitBuilder.WriteString("[white]Average Sale:[white]\n")
	if len(transactions) > 0 {
		avg := totalProfit / len(transactions)
		profitBuilder.WriteString(fmt.Sprintf("[yellow]  $%d[white]\n\n", avg))
	} else {
		profitBuilder.WriteString(fmt.Sprintf("[gray]  $0[white]\n\n"))
	}

	// Product breakdown
	profitBuilder.WriteString("\n[cyan]═══ INVENTORY ═══[white]\n\n")
	for _, prod := range products {
		profitBuilder.WriteString(fmt.Sprintf("[yellow]%s[white]\n", prod.Name))
		profitBuilder.WriteString(fmt.Sprintf("  [white]Stock: [cyan]%d[white]\n", prod.Stock))
		profitBuilder.WriteString(fmt.Sprintf("  [white]Price: [green]$%d[white]\n\n", prod.Price))
	}

	profitView.SetText(profitBuilder.String())
}

func buyIngredient() {
	if selectedIngredient >= len(ingredients) {
		return
	}

	ing := &ingredients[selectedIngredient]
	if userMoney >= ing.Price {
		userMoney -= ing.Price
		ing.Stock += 10
		updateBusinessViews()
		app.Draw()
	}
}

func prepareProduct() {
	if selectedProduct >= len(products) {
		return
	}

	prod := &products[selectedProduct]

	// Check if we have enough ingredients
	canPrepare := true
	for ingName, qtyNeeded := range prod.Ingredients {
		found := false
		for i := range ingredients {
			if ingredients[i].Name == ingName {
				if ingredients[i].Stock < qtyNeeded {
					canPrepare = false
				}
				found = true
				break
			}
		}
		if !found {
			canPrepare = false
		}
	}

	if canPrepare {
		// Deduct ingredients
		for ingName, qtyNeeded := range prod.Ingredients {
			for i := range ingredients {
				if ingredients[i].Name == ingName {
					ingredients[i].Stock -= qtyNeeded
					break
				}
			}
		}
		prod.Stock++
		updateBusinessViews()
		app.Draw()
	}
}

func deleteProduct() {
	if selectedProduct >= len(products) || len(products) <= 1 {
		return
	}

	products = append(products[:selectedProduct], products[selectedProduct+1:]...)
	if selectedProduct >= len(products) {
		selectedProduct = len(products) - 1
	}
	updateBusinessViews()
	app.Draw()
}

func showNewProductForm() {
	form := tview.NewForm().
		AddInputField("Product Name", "", 20, nil, nil).
		AddInputField("Price", "", 10, nil, nil).
		AddButton("Create", func() {
			nameField := form.GetFormItem(0).(*tview.InputField)
			priceField := form.GetFormItem(1).(*tview.InputField)

			name := nameField.GetText()
			var price int
			fmt.Sscanf(priceField.GetText(), "%d", &price)

			if name != "" && price > 0 {
				newProduct := Product{
					Name:        name,
					Ingredients: map[string]int{},
					Price:       price,
					Stock:       0,
				}
				products = append(products, newProduct)
				updateBusinessViews()
			}

			pages.RemovePage("form")
			app.Draw()
		}).
		AddButton("Cancel", func() {
			pages.RemovePage("form")
			app.Draw()
		})

	form.SetBorder(true).
		SetTitle(" [yellow]New Product[white] ").
		SetTitleAlign(tview.AlignCenter)

	pages.AddPage("form", modal(form, 50, 10), true, true)
}

func modal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func simulateMarketplace() {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Find products with stock
		availableProducts := []int{}
		for i, prod := range products {
			if prod.Stock > 0 {
				availableProducts = append(availableProducts, i)
			}
		}

		if len(availableProducts) == 0 {
			continue
		}

		// Random purchase
		prodIdx := availableProducts[rand.Intn(len(availableProducts))]
		prod := &products[prodIdx]
		amount := rand.Intn(min(3, prod.Stock)) + 1

		if prod.Stock >= amount {
			prod.Stock -= amount
			profit := prod.Price * amount
			totalProfit += profit
			userMoney += profit

			transaction := Transaction{
				BuyerName: names[rand.Intn(len(names))],
				Product:   prod.Name,
				Amount:    amount,
				Profit:    profit,
				Time:      time.Now(),
			}
			transactions = append(transactions, transaction)

			app.QueueUpdateDraw(func() {
				updateMarketplaceViews()
				if currentPage == 0 {
					updateBusinessViews()
				}
			})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
