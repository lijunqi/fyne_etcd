package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeToolBar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
	)
}
