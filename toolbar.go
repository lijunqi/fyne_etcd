package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeToolBar(w fyne.Window) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func ()  {
			hostName := widget.NewEntry()
			hostName.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "Please input hostname or IP.")
			port := widget.NewEntry()
			port.Validator = validation.NewRegexp(`^[0-9]+$`, "Please input port number.")
			remember := false
			hostForm := widget.NewFormItem("Hostname/IP", hostName)
			items := []*widget.FormItem{
				hostForm,
				widget.NewFormItem("Port", port),
				widget.NewFormItem("Remember me", widget.NewCheck("", func(checked bool) {
					remember = checked
				})),
			}

			newSessionFormDialog := dialog.NewForm("New session...", "Open", "Cancel", items, func(b bool) {
				if !b {
					return
				}
				var rememberText string
				if remember {
					rememberText = "and remember this login"
				}

				ETCDViewer.AddHost(hostName.Text)
				ETCDViewer.HostList.Refresh()
				ETCDViewer.Tabs.Refresh()

				/*
				etcdObj := EtcdObj{[]string{hostName}}
				data, err := etcdObj.ListAllV3()
				if err != nil {
					fmt.Printf("xxx Error: %v\n", err)
				} else {
					ETCDViewer.Tabs.BaseWidget.ExtendBaseWidget(ETCDViewer.Tabs)
					ETCDViewer.Tabs.Append(container.NewTabItem(hostname, makeTable(data)))
				}
				*/

				log.Println("Host:", hostName.Text, "Port:", port.Text, rememberText)
			}, w)
			newSessionFormDialog.Resize(fyne.NewSize(400, 200))
			newSessionFormDialog.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
	)
}
