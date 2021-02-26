package app

import (
	"fmt"
	"gotomate/app/automate"
	"gotomate/app/fiber"
	"log"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

var aw = new(automate.Window)
var newFiber = fiber.NewFiber
var buttons = automate.NewButtons

// CreateApp Initiate the app
func CreateApp() {

	aw.Menu = new(automate.Menu)
	aw.PrimaryListBoxModel = automate.NewAutomateModel()

	if err := (declarative.MainWindow{
		AssignTo:   &aw.MainWindow,
		Icon:       "icon.ico",
		Title:      "Gotomate",
		Background: declarative.SolidColorBrush{Color: walk.RGB(11, 11, 11)},
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
						Image:    "/new.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyN},
						OnTriggered: func() {
							buttons.CleanButtons()
							newFiber.CleanFiber()
						},
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Open,
						Text:        "Open",
						Image:       "/open.png",
						Enabled:     declarative.Bind("enabledCB.Checked"),
						Visible:     declarative.Bind("!openHiddenCB.Checked"),
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: func() { fmt.Println("Open a fiber") },
					},
					declarative.Menu{
						AssignTo: &aw.Menu.Folders,
						Image:    "/folder.png",
						Text:     "My Fibers",
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Save,
						Text:        "Save",
						Image:       "/save.png",
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyS},
						OnTriggered: func() { aw.SaveFiber(newFiber) },
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
						Image:    "/run.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyE},
						OnTriggered: func() {
							go newFiber.RunFiber()
						},
					},
					declarative.Action{
						AssignTo: &aw.Menu.Stop,
						Text:     "Stop",
						Image:    "/stop.png",
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
						AssignTo:        &aw.PrimaryListBox,
						Name:            "PrimaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						Model:           aw.PrimaryListBoxModel,
						OnItemActivated: aw.PlbItemActivated,
					},
					declarative.ListBox{
						AssignTo:        &aw.SecondaryListBox,
						Name:            "SecondaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						OnItemActivated: func() { aw.SlbItemActivated(newFiber) },
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
										TextColor: walk.Color(0xffffff),
									},
									declarative.TextEdit{
										AssignTo:      &aw.FiberNameInput,
										Alignment:     declarative.Alignment2D(walk.AlignHNearVCenter),
										Font:          declarative.Font{Family: "Roboto", PointSize: 9},
										CompactHeight: true,
										MaxSize:       declarative.Size{Width: 150},
									},
									declarative.Composite{
										Layout: declarative.VBox{},
										Children: []declarative.Widget{
											declarative.PushButton{
												AssignTo:  &aw.SaveButton,
												MaxSize:   declarative.Size{Width: 100},
												Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
												Text:      "Save",
												OnClicked: func() { aw.SaveFiber(newFiber) },
											},
										},
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

	aw.AddSavedFibersActions(newFiber)
	aw.MainWindow.Run()
}
