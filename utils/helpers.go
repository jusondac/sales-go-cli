package utils

import "github.com/rivo/tview"

// Min returns the minimum of two integers
func Min(a, b int) int {
if a < b {
return a
}
return b
}

// Modal creates a centered modal dialog
func Modal(p tview.Primitive, width, height int) tview.Primitive {
return tview.NewFlex().
AddItem(nil, 0, 1, false).
AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
AddItem(nil, 0, 1, false).
AddItem(p, height, 1, true).
AddItem(nil, 0, 1, false), width, 1, true).
AddItem(nil, 0, 1, false)
}
