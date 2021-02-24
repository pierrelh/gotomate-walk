package app

import (
	"encoding/json"
	"fmt"
	"gotomate/packages"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

var aw = new(AutomateWindow)
var fiber = new(Fiber)
var runningFiber = 0

// CreateApp Initiate the app
func CreateApp() {

	aw.Menu = new(AutomateMenu)
	aw.PrimaryListBoxModel = NewAutomateModel()

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
						AssignTo: &aw.Menu.Run,
						Text:     "Run",
						Image:    "/run.png",
						Shortcut: declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyR},
						OnTriggered: func() {
							runFiber()
						},
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Save,
						Text:        "Save",
						Image:       "/save.png",
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyS},
						OnTriggered: aw.saveFiber,
					},
					declarative.Action{
						AssignTo:    &aw.Menu.Exit,
						Text:        "Exit",
						OnTriggered: func() { aw.MainWindow.Close() },
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
						OnItemActivated: aw.plbItemActivated,
					},
					declarative.ListBox{
						AssignTo:        &aw.SecondaryListBox,
						Name:            "SecondaryList",
						Font:            declarative.Font{Family: "Roboto", PointSize: 9},
						MultiSelection:  false,
						OnItemActivated: aw.slbItemActivated,
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
											go runFiber()
										},
									},
									declarative.PushButton{
										AssignTo:  &aw.SaveButton,
										MaxSize:   declarative.Size{Width: 100},
										Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
										Text:      "Save",
										OnClicked: aw.saveFiber,
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

	AddSavedFibersActions()
	aw.MainWindow.Run()
}

// AutomateWindow Setting the automate window structure
type AutomateWindow struct {
	MainWindow            *walk.MainWindow
	Menu                  *AutomateMenu
	PrimaryListBox        *walk.ListBox
	PrimaryListBoxModel   *AutomateModel
	SecondaryListBox      *walk.ListBox
	SecondaryListBoxModel *AutomateModel
	FiberNameLabel        *walk.Label
	FiberNameInput        *walk.TextEdit
	RunButton             *walk.PushButton
	SaveButton            *walk.PushButton
	ScrollView            *walk.ScrollView
}

//AutomateMenu Setting the automate window's menu structure
type AutomateMenu struct {
	WindowMenu *walk.Menu
	Open       *walk.Action
	Folders    *walk.Menu
	Run        *walk.Action
	Save       *walk.Action
	Exit       *walk.Action
}

// AutomateModel Setting the model of automate ListBox
type AutomateModel struct {
	walk.ListModelBase
	items []AutomateItem
}

// AutomateItem Setting automate packages items
type AutomateItem struct {
	name  string
	value string
}

func (aw *AutomateWindow) saveFiber() {
	name := aw.FiberNameInput.Text()
	if name == "" {
		var dlg *walk.Dialog
		var acceptPB *walk.PushButton

		errDialog := declarative.Dialog{
			Icon:          "/icon.ico",
			Title:         "Error",
			AssignTo:      &dlg,
			DefaultButton: &acceptPB,
			MinSize: declarative.Size{
				Width:  200,
				Height: 150,
			},
			Layout: declarative.VBox{},
			Children: []declarative.Widget{
				declarative.Composite{
					Layout: declarative.VBox{},
					Children: []declarative.Widget{
						declarative.TextLabel{
							Text:      "No name given for the fiber",
							Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
							Font:      declarative.Font{Family: "Roboto", PointSize: 9},
						},
						declarative.PushButton{
							AssignTo: &acceptPB,
							Text:     "OK",
							Font:     declarative.Font{Family: "Roboto", PointSize: 9},
							OnClicked: func() {
								dlg.Cancel()
								aw.FiberNameInput.SetFocus()
							},
						},
					},
				},
			},
		}

		errDialog.Run(aw.MainWindow)
	} else {
		fiber.Name = name
		path := "saves/" + name + ".json"
		file, _ := json.Marshal(fiber)
		ioutil.WriteFile(path, file, 0644)
		AddSavedFibersActions()
	}
}

func (aw *AutomateWindow) plbItemActivated() {
	i := aw.PrimaryListBox.CurrentIndex()
	if i != -1 {
		item := &aw.PrimaryListBoxModel.items[i]
		value := item.name
		newModel := NewAutomateSubModel(value)
		aw.SecondaryListBox.SetModel(newModel)
		aw.SecondaryListBoxModel = newModel
	}
}

func (aw *AutomateWindow) slbItemActivated() {
	i := aw.SecondaryListBox.CurrentIndex()
	if i != -1 {
		plbIndex := aw.PrimaryListBox.CurrentIndex()
		plbItem := &aw.PrimaryListBoxModel.items[plbIndex]
		packageName := plbItem.name

		item := &aw.SecondaryListBoxModel.items[i]
		funcName := item.name

		data, dialog := CreateNewDialog(funcName)

		newInstruction := &FiberInstruction{
			Package:  packageName,
			FuncName: funcName,
			Data:     data,
		}
		CreateFiberButton(aw.ScrollView, funcName, packageName, dialog)
		fiber.Instructions = append(fiber.Instructions, newInstruction)
	}
}

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *AutomateModel {
	env := packages.Packages

	m := &AutomateModel{items: make([]AutomateItem, len(env))}

	for i, e := range env {
		name := strconv.Itoa(i)
		value := e

		m.items[i] = AutomateItem{value, name}
	}

	return m
}

// NewAutomateSubModel Gettings the automate's subpackages
func NewAutomateSubModel(key string) *AutomateModel {
	env := packages.SubPackages

	m := &AutomateModel{items: make([]AutomateItem, len(env[key]))}

	for i, e := range env[key] {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := strconv.Itoa(i)
		value := e

		m.items[i] = AutomateItem{value, name}
	}

	return m
}

// ItemCount return the length of an AutomateModel items
func (m *AutomateModel) ItemCount() int {
	return len(m.items)
}

// Value return the value of an AutomateModel item
func (m *AutomateModel) Value(index int) interface{} {
	return m.items[index].name
}
