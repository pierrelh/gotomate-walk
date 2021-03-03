package app

import (
	"gotomate/app/automate"
	"gotomate/app/automate/listbox"
	"gotomate/app/automate/menu"
	"gotomate/app/fiber"
	"log"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

var aw = &automate.Window{
	Menu:             new(menu.Menu),
	PrimaryListBox:   &listbox.ListBox{Model: listbox.NewPrimaryListModel()},
	SecondaryListBox: new(listbox.ListBox),
}

var newFiber = fiber.NewFiber

// CreateApp Initiate the app
func CreateApp() {

	if err := (declarative.MainWindow{
		AssignTo:   &aw.MainWindow,
		Icon:       "icon.ico",
		Title:      "Gotomate",
		Background: declarative.SolidColorBrush{Color: walk.RGB(106, 215, 229)},
		MinSize:    declarative.Size{Width: 320, Height: 240},
		Size:       declarative.Size{Width: 800, Height: 600},
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		MenuItems: []declarative.MenuItem{
			declarative.Menu{
				AssignTo: &aw.Menu.WindowMenu,
				Text:     "&File",
				Items: []declarative.MenuItem{
					declarative.Action{
						AssignTo: &aw.Menu.New,
						Text:     "New",
						Image:    "/menu-icons/new.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyN},
						OnTriggered: func() {
							aw.CreateNewFiber()
						},
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Import,
						Text:        "Import",
						Image:       "/menu-icons/import.png",
						Enabled:     declarative.Bind("enabledCB.Checked"),
						Visible:     declarative.Bind("!openHiddenCB.Checked"),
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: func() { aw.InitImportFiber() },
					},
					declarative.Menu{
						AssignTo: &aw.Menu.Folders,
						Image:    "/menu-icons/folder.png",
						Text:     "My Fibers",
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Save,
						Text:        "Save",
						Image:       "/menu-icons/save.png",
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyS},
						OnTriggered: func() { aw.InitSaveFiber() },
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Export,
						Text:        "Export",
						Image:       "/menu-icons/export.png",
						Enabled:     declarative.Bind("enabledCB.Checked"),
						Visible:     declarative.Bind("!openHiddenCB.Checked"),
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: func() { aw.ExportFiber() },
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Exit,
						Text:        "Exit",
						OnTriggered: func() { aw.MainWindow.Close() },
					},
				},
			},
			declarative.Menu{
				AssignTo: &aw.Menu.WindowMenu,
				Text:     "&Fiber",
				Items: []declarative.MenuItem{
					declarative.Action{
						AssignTo: &aw.Menu.Run,
						Text:     "Run",
						Image:    "/menu-icons/run.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyE},
						OnTriggered: func() {
							go newFiber.RunFiber()
						},
					},
					declarative.Action{
						AssignTo: &aw.Menu.Stop,
						Text:     "Stop",
						Image:    "/menu-icons/stop.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyQ},
						OnTriggered: func() {
							go newFiber.StopFiber()
						},
					},
				},
			},
		},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout:  declarative.HBox{MarginsZero: true, SpacingZero: true},
				MaxSize: declarative.Size{Height: 120},
				Children: []declarative.Widget{
					declarative.ListBox{
						AssignTo:        &aw.PrimaryListBox.ListBox,
						Name:            "PrimaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						Model:           aw.PrimaryListBox.Model,
						OnItemActivated: aw.PlbItemActivated,
					},
					declarative.ListBox{
						AssignTo:       &aw.SecondaryListBox.ListBox,
						Name:           "SecondaryList",
						Font:           declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection: false,
						OnItemActivated: func() {
							aw.SlbItemActivated(newFiber)
						},
					},
					declarative.Composite{
						Layout: declarative.VBox{},
						Children: []declarative.Widget{
							declarative.Composite{
								Layout: declarative.HBox{},
								Children: []declarative.Widget{
									declarative.Label{
										AssignTo:  &aw.FiberNameLabel,
										Alignment: declarative.Alignment2D(walk.AlignHFarVCenter),
										Font:      declarative.Font{Family: "Roboto", PointSize: 12, Underline: true, Bold: true},
										Text:      "Fiber Name :",
									},
									declarative.TextEdit{
										AssignTo:      &aw.FiberNameInput,
										Alignment:     declarative.Alignment2D(walk.AlignHNearVCenter),
										Font:          declarative.Font{Family: "Roboto", PointSize: 9},
										CompactHeight: true,
										MaxSize:       declarative.Size{Width: 150},
									},
								},
							},
							declarative.Composite{
								Layout: declarative.HBox{},
								Children: []declarative.Widget{
									declarative.PushButton{
										AssignTo: &aw.RunButton,
										MaxSize:  declarative.Size{Width: 100},
										Font:     declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
										Text:     "RUN",
										OnClicked: func() {
											go newFiber.RunFiber()
										},
									},
									declarative.PushButton{
										AssignTo: &aw.StopButton,
										MaxSize:  declarative.Size{Width: 100},
										Font:     declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
										Text:     "STOP",
										OnClicked: func() {
											go newFiber.StopFiber()
										},
									},
								},
							},
						},
					},
				},
			},
			declarative.ScrollView{
				AssignTo:        &aw.ScrollView,
				Layout:          declarative.Flow{Margins: declarative.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Background:      declarative.SolidColorBrush{Color: walk.RGB(11, 11, 11)},
				HorizontalFixed: false,
				VerticalFixed:   false,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	aw.AddSavedFibersActions()
	aw.MainWindow.Run()
}
