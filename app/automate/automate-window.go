package automate

import (
	"gotomate/app/automate/list"
	"gotomate/app/fiber"
	"gotomate/packages"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

var pressed = false

//Window Setting the automate window structure
type Window struct {
	MainWindow            *walk.MainWindow
	Menu                  *Menu
	PrimaryListBox        *walk.ListBox
	PrimaryListBoxModel   *list.Model
	SecondaryListBox      *walk.ListBox
	SecondaryListBoxModel *list.Model
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
	New        *walk.Action
	Open       *walk.Action
	Folders    *walk.Menu
	Run        *walk.Action
	Stop       *walk.Action
	Save       *walk.Action
	Exit       *walk.Action
}

//PlbItemActivated Fill the Secondary list when an element of the Primary is activated
func (aw *Window) PlbItemActivated() {
	i := aw.PrimaryListBox.CurrentIndex()
	if i != -1 {
		item := &aw.PrimaryListBoxModel.Items[i]
		value := item.Name
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
		plbItem := &aw.PrimaryListBoxModel.Items[plbIndex]
		packageName := plbItem.Name

		item := &aw.SecondaryListBoxModel.Items[i]
		funcName := item.Name

		data, dialog := packages.CreateNewDialog(funcName)

		newInstruction := &fiber.Instruction{
			Package:  packageName,
			FuncName: funcName,
			Data:     data,
		}
		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog)
	}
}

// NewAutomateModel Getting the automates packages
func NewAutomateModel() *list.Model {
	env := packages.Packages

	m := &list.Model{Items: make([]list.Item, len(env))}

	for i, e := range env {
		value := strconv.Itoa(i)
		name := e

		m.Items[i] = list.Item{Name: name, Value: value}
	}

	return m
}

// NewAutomateSubModel Gettings the automate's subpackages
func NewAutomateSubModel(key string) *list.Model {
	env := packages.SubPackages

	m := &list.Model{Items: make([]list.Item, len(env[key]))}

	for i, e := range env[key] {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		value := strconv.Itoa(i)
		name := e

		m.Items[i] = list.Item{Name: name, Value: value}
	}

	return m
}

// CreateFiberButton Create a new Fiberbutton in the fiber
func (aw *Window) CreateFiberButton(instruction *fiber.Instruction, dialog *packages.Dialog) error {
	fb := new(Button)
	NewButtons.Buttons = append(NewButtons.Buttons, fb)
	fb.Dialog = dialog
	bmp, err := walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/" + instruction.Package + ".png")
	if err != nil {
		bmp, _ = walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/func-icons/default.png")
	}

	if err := (declarative.Composite{
		AssignTo:   &fb.Composite,
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		Background: declarative.BitmapBrush{Image: bmp},
		Alignment:  declarative.Alignment2D(walk.AlignHNearVCenter),
		OnMouseDown: func(x, y int, button walk.MouseButton) {
			pressed = true
		},
		OnMouseMove: func(x, y int, button walk.MouseButton) {
			pressed = false
		},
		OnMouseUp: func(x, y int, button walk.MouseButton) {
			if pressed {
				switch button {
				case 1:
					fb.Dialog.DialogContent.Run(aw.MainWindow)
					break
				case 2:
					fb.DeleteButton(aw, fiber.NewFiber)
					break
				default:
					break
				}
			}
		},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.LinkLabel{
				AssignTo:  &fb.LinkLabel,
				Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
				Text:      instruction.FuncName,
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
			},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(aw.ScrollView)); err != nil {
		return err
	}
	fb.Composite.SetCursor(walk.CursorHand())
	fb.Composite.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 150}, walk.Size{Width: 300, Height: 150})
	fb.LinkLabel.SetMinMaxSizePixels(walk.Size{Width: 300, Height: 20}, walk.Size{Width: 300, Height: 20})
	fb.Composite.Children().At(0).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 100}, walk.Size{Width: 300, Height: 100})
	fb.Composite.Children().At(2).SetMinMaxSizePixels(walk.Size{Width: 300, Height: 30}, walk.Size{Width: 300, Height: 30})
	return nil
}
