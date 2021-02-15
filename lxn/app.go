package lxn

import (
	"gotomate/automate"
	"log"
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

// CreateApp Initiate the app
func CreateApp() {
	aw := &AutomateWindow{}
	aw.plbmodel = NewAutomateModel()

	if _, err := (MainWindow{
		AssignTo: &aw.MainWindow,
		Icon:     "img/icon.ico",
		Title:    "Gotomate",
		// Background: SolidColorBrush{Color: walk.RGB(0, 0, 0)},
		MinSize: Size{Width: 320, Height: 240},
		Size:    Size{Width: 800, Height: 600},
		Layout:  VBox{MarginsZero: true},
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
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

// AutomateWindow Setting the automate window structure
type AutomateWindow struct {
	*walk.MainWindow
	plb      *walk.ListBox
	plbmodel *AutomateModel
	slb      *walk.ListBox
	slbmodel *AutomateModel
	tw       *walk.TabWidget
	te       *walk.TextEdit
	pb       *walk.PushButton
	sv       *walk.ScrollView
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
	i := aw.slb.CurrentIndex()
	item := &aw.slbmodel.items[i]
	value := item.name
	a, _ := walk.NewTextEdit(aw.sv)
	a.SetCompactHeight(true)
	a.SetTextAlignment(walk.AlignCenter)
	a.SetText(value)
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

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *AutomateModel {
	env := automate.Packages

	m := &AutomateModel{items: make([]AutomateItem, len(env))}

	for i, e := range env {
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

// NewAutomateSubModel Gettings the automate's subpackages
func NewAutomateSubModel(key string) *AutomateModel {
	env := automate.SubPackages

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

func (m *AutomateModel) ItemCount() int {
	return len(m.items)
}

func (m *AutomateModel) Value(index int) interface{} {
	return m.items[index].name
}
