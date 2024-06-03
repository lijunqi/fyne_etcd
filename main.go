package main

import (
	"fyne_etcd/etcdviewer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

var ETCDViewer = etcdviewer.EtcdViewer{}



func main() {
	a := app.New()
	w := a.NewWindow("FyneETCD")

	toolbar := makeToolBar(w)

	left := ETCDViewer.MakeHostList()
	//kvTable := makeTable(&w)
	right := ETCDViewer.MakeAppTabs(w)

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