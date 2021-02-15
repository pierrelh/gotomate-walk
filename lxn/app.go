package lxn

import (
	"log"
	"piproto/automate"
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
	mw := &AutomateWindow{model: NewAutomateModel()}

	if _, err := (MainWindow{
		AssignTo:   &mw.MainWindow,
		Icon:       "img/icon.ico",
		Title:      "Gotomate",
		Background: SolidColorBrush{Color: walk.RGB(0, 0, 0)},
		MinSize:    Size{240, 320},
		Size:       Size{600, 800},
		Layout:     VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &mw.te,
						ReadOnly: true,
					},
					ListBox{
						Font:                  Font{Family: "Segoe UI", PointSize: 9},
						AssignTo:              &mw.slb,
						OnCurrentIndexChanged: mw.slbItemActivated,
						OnItemActivated:       mw.slbItemActivated,
					},
					ListBox{
						Font:                  Font{Family: "Segoe UI", PointSize: 9},
						AssignTo:              &mw.plb,
						Model:                 mw.model,
						OnCurrentIndexChanged: mw.plbItemActivated,
						OnItemActivated:       mw.plbItemActivated,
					},
				},
			},
			PushButton{
				AssignTo:  &mw.pb,
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
	model *AutomateModel
	plb   *walk.ListBox
	slb   *walk.ListBox
	te    *walk.TextEdit
	pb    *walk.PushButton
}

func (mw *AutomateWindow) plbItemActivated() {
	i := mw.plb.CurrentIndex()
	item := &mw.model.items[i]
	value := item.name
	mw.slb.SetModel(NewAutomateSubModel(value))
}

func (mw *AutomateWindow) slbItemActivated() {
	// value := mw.model.items[mw.slb.CurrentIndex()].value

	// walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
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

	m := &AutomateModel{items: make([]AutomateItem, len(env))}

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
