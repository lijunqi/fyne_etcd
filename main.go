package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var dataKey = []string{"localhost", "ItemB", "ItemC", "ItemD"}
var dataSet = map[string] [][]string {
	"localhost": {
		{"key A1", "value A1"},
	},
	"ItemB": {
		{"key B1", "value B1"},
		{"key B2", "value B2"},
	},
	"ItemC": {
		{"key C1", "value C1"},
		{"key C2", "value C2"},
		{"key C3", "value C3"},
	},
	"ItemD": {
		{"key D1", "value D1"},
		{"key D2", "value D2"},
		{"key D3", "value D3"},
		{"key D4", "value D4"},
	},
}


func makeList() *widget.List {

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	list := widget.NewList(
		func() int {
			return len(dataKey)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), canvas.NewText(dataKey[0], color.Black))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*fyne.Container).Objects[0] = canvas.NewText("Icon", color.Black)
			o.(*fyne.Container).Objects[1] = canvas.NewText(dataKey[i], blue)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(dataKey[id])
		fmt.Printf("Select %s\n", dataKey[id])
		icon.SetResource(theme.ComputerIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
		fmt.Printf("Unselect %s\n", dataKey[id])
	}
	list.Select(0)

	return list
}

func makeAppTabsTab(_ fyne.Window) fyne.CanvasObject {
	tabs := &container.AppTabs{}
	tabs.BaseWidget.ExtendBaseWidget(tabs)
	for _, v := range dataKey {
		tabs.Append(container.NewTabItem(v, makeTable(nil, dataSet[v])))
	}
	return tabs

	/*
	return container.NewAppTabs(
		container.NewTabItem("Tab 1", makeTable(nil)),
		container.NewTabItem("Tab 2 bigger", widget.NewLabel("Content of tab 2")),
		container.NewTabItem("Tab 3", widget.NewLabel("Content of tab 3")),
	)
	*/
}


func makeTable(w *fyne.Window, data [][]string) *widget.Table {

	//etcdObj := etcdcli.EtcdObj{}
	//data, err := etcdObj.ListAllV3()
	//if err != nil {
	//	fmt.Printf("xxx Error: %v\n", err)
	//	dialog.ShowError(err, *w)
	//	return nil
	//}


	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])+1 // +1 for 1st col(sequence number col)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				cell.(*widget.Label).SetText(data[id.Row][0])
			case 2:
				cell.(*widget.Label).SetText(data[id.Row][1])
			default:
				cell.(*widget.Label).SetText(data[id.Row][1])
			}
		},
	)
	
	table.SetColumnWidth(0, 30)  // Set 1st column width
	table.SetColumnWidth(1, 200) // Set 2nd ~ last columns width

	return table
}

func main() {
	a := app.New()
	w := a.NewWindow("FyneETCD")

	toolbar := makeToolBar()

	left := makeList()
	//kvTable := makeTable(&w)
	right := makeAppTabsTab(w)

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {}),
		widget.NewButton("Light", func() {}),
	)

	leftBorder := container.NewBorder(nil, themes, nil, nil, left)
	rightBorder := container.NewBorder(nil, nil, nil, nil, right)

	content := container.NewHSplit(leftBorder, rightBorder)
	content.SetOffset(0.25)

	mainContent := container.NewBorder(toolbar, nil, nil, nil, content)
	w.SetContent(mainContent)

	w.Resize(fyne.NewSize(640, 480))
	w.SetMaster()
	w.CenterOnScreen()
	w.ShowAndRun()
}