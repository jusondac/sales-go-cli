package models

import "github.com/rivo/tview"

// AppState holds all global application state
type AppState struct {
	App         *tview.Application
	Pages       *tview.Pages
	CurrentPage int // 0: Business, 1: Marketplace

	// Business panel references
	IngredientsListView *tview.TextView
	PreparationView     *tview.TextView
	InfoView            *tview.TextView

	// Marketplace panel references
	TransactionsView *tview.TextView
	ProfitView       *tview.TextView

	// Data
	Ingredients        []Ingredient
	Products           []Product
	Transactions       []Transaction
	SelectedIngredient int
	SelectedProduct    int
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
		{Name: "Flour", Price: 10, Stock: 0},
		{Name: "Sugar", Price: 15, Stock: 0},
		{Name: "Eggs", Price: 20, Stock: 0},
		{Name: "Butter", Price: 25, Stock: 0},
		{Name: "Milk", Price: 12, Stock: 0},
		{Name: "Chocolate", Price: 30, Stock: 0},
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

	s.Transactions = []Transaction{}
}
