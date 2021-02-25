package automate

import (
	"gotomate/app/fiber"
	"gotomate/packages"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//Window Setting the automate window structure
type Window struct {
	MainWindow            *walk.MainWindow
	Menu                  *Menu
	PrimaryListBox        *walk.ListBox
	PrimaryListBoxModel   *ListModel
	SecondaryListBox      *walk.ListBox
	SecondaryListBoxModel *ListModel
	FiberNameLabel        *walk.Label
	FiberNameInput        *walk.TextEdit
	RunButton             *walk.PushButton
	StopButton            *walk.PushButton
	SaveButton            *walk.PushButton
	ScrollView            *walk.ScrollView
}

//Menu Setting the automate window's menu structure
type Menu struct {
	WindowMenu *walk.Menu
	Open       *walk.Action
	Folders    *walk.Menu
	Run        *walk.Action
	Stop       *walk.Action
	Save       *walk.Action
	Exit       *walk.Action
}

//ListModel Setting the model of automate ListBox
type ListModel struct {
	walk.ListModelBase
	items []ListItem
}

//ListItem Setting automate packages items
type ListItem struct {
	name  string
	value string
}

//PlbItemActivated Fill the Secondary list when an element of the Primary is activated
func (aw *Window) PlbItemActivated() {
	i := aw.PrimaryListBox.CurrentIndex()
	if i != -1 {
		item := &aw.PrimaryListBoxModel.items[i]
		value := item.name
		newModel := NewAutomateSubModel(value)
		aw.SecondaryListBox.SetModel(newModel)
		aw.SecondaryListBoxModel = newModel
	}
}

//SlbItemActivated Create a fiber's button when an element of the Secondary list is activated
func (aw *Window) SlbItemActivated(currentFiber *fiber.Fiber) {
	i := aw.SecondaryListBox.CurrentIndex()
	if i != -1 {
		plbIndex := aw.PrimaryListBox.CurrentIndex()
		plbItem := &aw.PrimaryListBoxModel.items[plbIndex]
		packageName := plbItem.name

		item := &aw.SecondaryListBoxModel.items[i]
		funcName := item.name

		data, dialog := packages.CreateNewDialog(funcName)

		newInstruction := &fiber.Instruction{
			Package:  packageName,
			FuncName: funcName,
			Data:     data,
		}
		aw.CreateFiberButton(aw.ScrollView, funcName, packageName, dialog)
		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
	}
}

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *ListModel {
	env := packages.Packages

	m := &ListModel{items: make([]ListItem, len(env))}

	for i, e := range env {
		name := strconv.Itoa(i)
		value := e

		m.items[i] = ListItem{value, name}
	}

	return m
}

// NewAutomateSubModel Gettings the automate's subpackages
func NewAutomateSubModel(key string) *ListModel {
	env := packages.SubPackages

	m := &ListModel{items: make([]ListItem, len(env[key]))}

	for i, e := range env[key] {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := strconv.Itoa(i)
		value := e

		m.items[i] = ListItem{value, name}
	}

	return m
}

// ItemCount return the length of an AutomateModel items
func (m *ListModel) ItemCount() int {
	return len(m.items)
}

// Value return the value of an AutomateModel item
func (m *ListModel) Value(index int) interface{} {
	return m.items[index].name
}

// CreateFiberButton Create a new Fiberbutton in the fiber
func (aw *Window) CreateFiberButton(parent *walk.ScrollView, funcName, packageName string, dialog declarative.Dialog) error {
	fb := new(fiber.Button)
	fb.DialogWindow = dialog
	bmp, err := walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/" + packageName + ".png")
	if err != nil {
		bmp, _ = walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/default.png")
	}

	if err := (declarative.Composite{
		AssignTo:   &fb.Composite,
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		Background: declarative.BitmapBrush{Image: bmp},
		Alignment:  declarative.Alignment2D(walk.AlignHNearVCenter),
		OnMouseDown: func(x, y int, button walk.MouseButton) {
			fb.DialogWindow.Run(aw.MainWindow)
		},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.LinkLabel{
				AssignTo:  &fb.LinkLabel,
				Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
				Text:      funcName,
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
			},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(parent)); err != nil {
		return err
	}
	fb.Composite.SetCursor(walk.CursorHand())
	fb.Composite.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 150}, walk.Size{Width: 300, Height: 150})
	fb.LinkLabel.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 20}, walk.Size{Width: 300, Height: 20})
	fb.Composite.Children().At(0).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 100}, walk.Size{Width: 300, Height: 100})
	fb.Composite.Children().At(2).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 30}, walk.Size{Width: 300, Height: 30})
	return nil
}
