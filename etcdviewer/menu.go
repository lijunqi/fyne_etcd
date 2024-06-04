package etcdviewer

import (
	"fyne.io/fyne/v2"
)

func (v *EtcdViewer) MakeMenu() *fyne.MainMenu {
	// create three menu items
	openMenuItem := fyne.NewMenuItem("Open...", func(){})
	saveMenuItem := fyne.NewMenuItem("Save", func(){})
	saveAsMenuItem := fyne.NewMenuItem("Save as...", func(){})

	// create a file menu, and add the three items to it
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// create a main menu, and add the file menu to it
	v.menu = fyne.NewMainMenu(fileMenu)
	return v.menu
}
