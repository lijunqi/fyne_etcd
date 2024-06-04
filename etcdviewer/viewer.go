package etcdviewer

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type EtcdViewer struct {
	menu     *fyne.MainMenu
	toolbar  *widget.Toolbar
	Tabs     *container.AppTabs
	HostList *widget.List
	hostNames []string
}

func (v *EtcdViewer) MakeHostList() *widget.List {

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	v.HostList = widget.NewList(
		func() int {
			return len(v.hostNames)
		},
		func() fyne.CanvasObject {
			if len(v.hostNames) > 0 {
				return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), canvas.NewText(v.hostNames[0], color.Black))
			} else {
				return container.NewHBox(widget.NewIcon(theme.ComputerIcon()), canvas.NewText("No host", color.Black))
			}
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*fyne.Container).Objects[0] = canvas.NewText("Icon", color.Black)
			o.(*fyne.Container).Objects[1] = canvas.NewText(v.hostNames[i], blue)
		},
	)
	v.HostList.OnSelected = func(id widget.ListItemID) {
		label.SetText(v.hostNames[id])
		fmt.Printf("Select %s\n", v.hostNames[id])
		icon.SetResource(theme.ComputerIcon())
	}
	v.HostList.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
		fmt.Printf("Unselect %s\n", v.hostNames[id])
	}
	v.HostList.Select(0)

	return v.HostList
}

func (v *EtcdViewer) AddHost(hostname string) {
	v.hostNames = append(v.hostNames, hostname)
}


func (v *EtcdViewer) MakeAppTabs(_ fyne.Window) fyne.CanvasObject {
	if len(v.hostNames) == 0 {
		return v.Tabs
	}

	etcdObj := EtcdObj{[]string{v.hostNames[0]}}
	data, err := etcdObj.ListAllV3()
	if err != nil {
		fmt.Printf("xxx Error: %v\n", err)
		//dialog.ShowError(err, *w)
		return v.Tabs
	}

	for _, hostname := range v.hostNames {
		v.Tabs.Append(container.NewTabItem(hostname, makeTable(data)))
	}
	return v.Tabs

	/*
	return container.NewAppTabs(
		container.NewTabItem("Tab 1", makeTable(nil)),
		container.NewTabItem("Tab 2 bigger", widget.NewLabel("Content of tab 2")),
		container.NewTabItem("Tab 3", widget.NewLabel("Content of tab 3")),
	)
	*/
}

func makeTable(data [][]string) *widget.Table {
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

func (v *EtcdViewer) MakeToolBar(w fyne.Window) *widget.Toolbar {
	v.toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func ()  {
			hostName := widget.NewEntry()
			hostName.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "Please input hostname or IP.")
			hostName.Text = "localhost"
			port := widget.NewEntry()
			port.Validator = validation.NewRegexp(`^[0-9]+$`, "Please input port number.")
			port.Text = "2379"
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

				v.AddHost(hostName.Text)
				v.HostList.Refresh()

				endpoint := fmt.Sprintf("%s:%s", hostName.Text, port.Text)
				etcdObj := EtcdObj{[]string{endpoint}}
				data, err := etcdObj.ListAllV3()
				if err != nil {
					fmt.Printf("xxx Error: %v\n", err)
				} else {
					v.Tabs.Append(container.NewTabItem(hostName.Text, makeTable(data)))
					v.Tabs.Refresh()
					fmt.Println("Refresh Tabs")
				}

				log.Println("Host:", hostName.Text, "Port:", port.Text, rememberText)
			}, w)
			newSessionFormDialog.Resize(fyne.NewSize(400, 200))
			newSessionFormDialog.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
	)
	return v.toolbar
}
