package models

import "github.com/rivo/tview"

// AppState holds all global application state
type AppState struct {
	App         *tview.Application
	Pages       *tview.Pages
	CurrentPage int // 0: Business, 1: Marketplace, 2: Analytics

	// Business panel references
	IngredientsListView *tview.TextView
	PreparationView     *tview.TextView
	ProductInfoView     *tview.TextView
	InfoView            *tview.TextView

	// Marketplace panel references
	TransactionsView *tview.TextView
	ProfitView       *tview.TextView
	TaxesList        *tview.List

	// Analytics panel references
	IngredientPricesView *tview.TextView
	ProductPricesView    *tview.TextView
	PriceHistoryView     *tview.TextView

	// Top status bar
	StatusBarView *tview.TextView

	// Bottom help panel
	HelpPanelView *tview.TextView

	// Data
	Ingredients        []Ingredient
	Products           []Product
	Transactions       []Transaction
	Taxes              []Tax
	SelectedIngredient int
	SelectedProduct    int
	SelectedTax        int
	TotalProfit        int
	UserMoney          int
}

// NewAppState creates and initializes a new application state
func NewAppState() *AppState {
	state := &AppState{
		CurrentPage: 0,
		UserMoney:   1000,
	}
	state.InitializeData()
	return state
}

// InitializeData sets up initial ingredients and products
func (s *AppState) InitializeData() {
	s.Ingredients = []Ingredient{
		{Name: "Flour", Price: 10, Stock: 0, Step: 2, Floor: 5, Ceil: 20},
		{Name: "Sugar", Price: 15, Stock: 0, Step: 3, Floor: 8, Ceil: 25},
		{Name: "Eggs", Price: 20, Stock: 0, Step: 4, Floor: 10, Ceil: 35},
		{Name: "Butter", Price: 25, Stock: 0, Step: 5, Floor: 15, Ceil: 40},
		{Name: "Milk", Price: 12, Stock: 0, Step: 2, Floor: 6, Ceil: 22},
		{Name: "Chocolate", Price: 30, Stock: 0, Step: 6, Floor: 20, Ceil: 50},
	}

	s.Products = []Product{
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
			Step:  10,
			Floor: 70,
			Ceil:  150,
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
			Step:  5,
			Floor: 35,
			Ceil:  80,
		},
	}

	s.Transactions = []Transaction{}
}
