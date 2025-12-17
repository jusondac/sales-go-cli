# DK Business Simulator

A terminal-based business simulation game written in Go. Manage your inventory, craft products, and trade in a dynamic marketplace, all from your command line.

## Features

- **Business Management**: Buy raw ingredients and craft them into sellable products.
- **Dynamic Marketplace**: Real-time price simulation for ingredients and products.
- **Analytics**: Track your financial performance and inventory stats.
- **Rich TUI**: Interactive terminal user interface built with `tview`.

## Installation & Run

Ensure you have [Go](https://go.dev/) installed (1.25+ recommended).

```bash
# Clone the repository
git clone <repository-url>
cd sales-go-cli

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

## Controls

### Global Navigation
- **Left / Right Arrow Keys**: Switch between Business, Marketplace, and Analytics panels.
- **Ctrl+C**: Quit the application.

### Business Panel
- **`w` / `s`**: Select previous/next ingredient.
- **`e` / `d`**: Select previous/next product.
- **`b`**: Buy selected ingredient.
- **`p`**: Prepare (craft) selected product.
- **`n`**: Create a new product type.
- **`x`**: Delete selected product type.

## Project Structure

- **`main.go`**: Entry point and application initialization.
- **`handlers/`**: Game logic, input handling, and simulation routines.
- **`models/`**: Data structures for state, products, and ingredients.
- **`panels/`**: UI layout and component definitions for each screen.
- **`utils/`**: Helper functions.
