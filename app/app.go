package app

import (
	"gotomate/automate"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

/*
   Use:
   Databinding
   listbox
   log view
   multipage
   notifyicon
*/
var aw = &AutomateWindow{}
var fiber = &Fiber{}

// CreateApp Initiate the app
func CreateApp() {

	aw.plbmodel = NewAutomateModel()

	window := MainWindow{
		AssignTo:   &aw.mw,
		Icon:       "img/icon.ico",
		Title:      "Gotomate",
		Background: SolidColorBrush{Color: walk.RGB(0, 0, 0)},
		MinSize:    Size{Width: 320, Height: 240},
		Size:       Size{Width: 800, Height: 600},
		Layout:     VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				MaxSize: Size{Width: 600, Height: 100},
				Children: []Widget{
					ListBox{
						Name:            "PrimaryList",
						Font:            Font{Family: "Segoe UI", PointSize: 9},
						MaxSize:         Size{Width: 600, Height: 100},
						AssignTo:        &aw.plb,
						Model:           aw.plbmodel,
						OnItemActivated: aw.plbItemActivated,
					},
					ListBox{
						Name:            "SecondaryList",
						Font:            Font{Family: "Segoe UI", PointSize: 9},
						MaxSize:         Size{Width: 600, Height: 100},
						AssignTo:        &aw.slb,
						OnItemActivated: aw.slbItemActivated,
						BindingMember:   "test",
						DisplayMember:   "test",
					},
				},
			},
			ScrollView{
				AssignTo:        &aw.sv,
				MinSize:         Size{Width: 600, Height: 200},
				MaxSize:         Size{Width: 600, Height: 800},
				HorizontalFixed: false,
				VerticalFixed:   false,
				Layout:          VBox{MarginsZero: true},
			},
			PushButton{
				AssignTo:  &aw.pb,
				Text:      "RUN",
				OnClicked: automate.LaunchAutomate,
			},
		},
	}
	window.Run()
}

// AutomateWindow Setting the automate window structure
type AutomateWindow struct {
	mw         *walk.MainWindow
	plb        *walk.ListBox
	plbmodel   *AutomateModel
	slb        *walk.ListBox
	slbmodel   *AutomateModel
	pb         *walk.PushButton
	sv         *walk.ScrollView
	pushButton *walk.PushButton
}

// AutomateItem Setting automate packages items
type AutomateItem struct {
	name  string
	value string
}

// AutomateModel Setting the model of automate ListBox
type AutomateModel struct {
	walk.ListModelBase
	items []AutomateItem
}

func (aw *AutomateWindow) plbItemActivated() {
	i := aw.plb.CurrentIndex()
	item := &aw.plbmodel.items[i]
	value := item.name
	newModel := NewAutomateSubModel(value)
	aw.slb.SetModel(newModel)
	aw.slbmodel = newModel
}

func (aw *AutomateWindow) slbItemActivated() {
	plbIndex := aw.plb.CurrentIndex()
	plbItem := &aw.plbmodel.items[plbIndex]
	packageName := plbItem.name

	i := aw.slb.CurrentIndex()
	item := &aw.slbmodel.items[i]
	funcName := item.name

	fiberButton, _ := CreatePushButton(aw.sv, funcName)

	newInstruction := &FiberInstruction{
		Package:  packageName,
		FuncName: funcName,
		Button:   fiberButton,
	}
	fiber.Instructions = append(fiber.Instructions, newInstruction)
}

// CreatePushButton Create a new push button in the fiber
func CreatePushButton(parent walk.Container, funcName string) (*FiberButton, error) {
	pb, err := walk.NewPushButton(parent)
	if err != nil {
		return nil, err
	}

	pb.SetText(funcName)

	fiberButton := &FiberButton{
		pb,
		CreateNewDialog(funcName),
	}

	if err := walk.InitWrapperWindow(fiberButton); err != nil {
		return nil, err
	}

	return fiberButton, nil
}

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *AutomateModel {
	env := Packages

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
	env := SubPackages

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
