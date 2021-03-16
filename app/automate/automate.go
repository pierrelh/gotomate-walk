package automate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotomate/app/automate/button"
	"gotomate/app/automate/dialogs"
	"gotomate/app/automate/listbox"
	"gotomate/app/automate/menu"
	"gotomate/fiber"
	"gotomate/fiber/packages"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//NewButtons Create the automate's list of buttons
var NewButtons = button.NewButtons
var currentFiber = fiber.NewFiber
var pressed = false
var moved = false
var clientX = 0
var clientY = 0
var row = 1

//Window Setting the automate window structure
type Window struct {
	MainWindow       *walk.MainWindow
	Menu             *menu.Menu
	PrimaryListBox   *listbox.ListBox
	SecondaryListBox *listbox.ListBox
	FiberNameLabel   *walk.Label
	FiberNameInput   *walk.TextEdit
	RunButton        *walk.PushButton
	StopButton       *walk.PushButton
	ScrollView       *walk.ScrollView
}

//PlbItemActivated Fill the Secondary list when an element of the Primary is activated
func (aw *Window) PlbItemActivated() {
	i := aw.PrimaryListBox.ListBox.CurrentIndex()
	if i != -1 {
		item := &aw.PrimaryListBox.Model.Items[i]
		value := item.Name
		newModel := aw.PrimaryListBox.NewSecondaryListModel(value)
		aw.SecondaryListBox.ListBox.SetModel(newModel)
		aw.SecondaryListBox.Model = newModel
	}
}

//SlbItemActivated Create a fiber's button when an element of the Secondary list is activated
func (aw *Window) SlbItemActivated(currentFiber *fiber.Fiber) {
	i := aw.SecondaryListBox.ListBox.CurrentIndex()
	if i != -1 {
		plbIndex := aw.PrimaryListBox.ListBox.CurrentIndex()
		plbItem := &aw.PrimaryListBox.Model.Items[plbIndex]
		packageName := plbItem.Name

		item := &aw.SecondaryListBox.Model.Items[i]
		funcName := item.Name

		instructionId := 1
		if len(currentFiber.Instructions) > 1 {
			instructionId = currentFiber.Instructions[len(currentFiber.Instructions)-1].ID + 1
		}
		nextId := instructionId + 1

		data, dialog := packages.CreateNewDialog(packageName, funcName)

		newInstruction := &fiber.Instruction{
			ID:                instructionId,
			Package:           packageName,
			FuncName:          funcName,
			X:                 0,
			Y:                 0,
			NextInstructionID: nextId,
			InstructionData:   data,
		}
		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog, true)
	}
}

// CreateFiberButton Create a new Fiberbutton in the fiber
func (aw *Window) CreateFiberButton(instruction *fiber.Instruction, dialog *packages.Dialog, isDeletable bool) error {
	fb := new(button.Button)
	NewButtons.Buttons = append(NewButtons.Buttons, fb)
	fb.Dialog = dialog
	bmp, err := walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/fiber/packages/" + instruction.Package + "/icon.png")
	if err != nil {
		bmp, _ = walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/img/default.png")
	}

	if err := (declarative.Composite{
		AssignTo:   &fb.Composite,
		Layout:     declarative.Grid{Columns: 1, Margins: declarative.Margins{Left: 10, Top: 10, Right: 10, Bottom: 5}, SpacingZero: true},
		MinSize:    declarative.Size{Width: 300, Height: 150},
		MaxSize:    declarative.Size{Width: 300, Height: 150},
		Row:        row,
		Background: declarative.BitmapBrush{Image: bmp},
		OnMouseDown: func(x, y int, button walk.MouseButton) {
			pressed = true
			clientX, clientY = robotgo.GetMousePos()
		},
		OnMouseMove: func(x, y int, button walk.MouseButton) {
			if pressed {
				mx, my := robotgo.GetMousePos()
				difX := clientX - mx
				difY := clientY - my
				clientX = mx
				clientY = my
				moved = true
				if fb.Composite.XPixels()-difX < 0 {
					fb.Composite.SetXPixels(0)
					instruction.X = 0
				} else if (fb.Composite.X() - difX) > aw.ScrollView.Layout().Margins().HFar {
					fb.Composite.SetX(aw.ScrollView.Layout().Margins().HFar)
					instruction.X = aw.ScrollView.Layout().Margins().HFar
				} else {
					fb.Composite.SetXPixels(fb.Composite.XPixels() - difX)
					instruction.X = fb.Composite.XPixels() - difX
				}
				if fb.Composite.YPixels()-difY < 0 {
					fb.Composite.SetYPixels(0)
					instruction.Y = 0
				} else if (fb.Composite.Y() - difY) > aw.ScrollView.Layout().Margins().VFar {
					fb.Composite.SetY(aw.ScrollView.Layout().Margins().VFar)
					instruction.Y = aw.ScrollView.Layout().Margins().VFar
				} else {
					fb.Composite.SetYPixels(fb.Composite.YPixels() - difY)
					instruction.Y = fb.Composite.YPixels() - difY
				}
			}
		},
		OnMouseUp: func(x, y int, button walk.MouseButton) {
			if pressed && !moved {
				switch button {
				case 1:
					fb.Dialog.DialogContent.Run(aw.MainWindow)
					break
				case 2:
					if isDeletable {
						aw.DeleteFiberButton(fb)
					}
					break
				default:
					break
				}
			}
			pressed = false
			moved = false
		},
		Children: []declarative.Widget{
			declarative.LinkLabel{
				AssignTo:  &fb.IDLabel,
				MinSize:   declarative.Size{Height: 50},
				MaxSize:   declarative.Size{Height: 50},
				Font:      declarative.Font{Family: "Roboto", PointSize: 7},
				Text:      strconv.Itoa(instruction.ID),
				Alignment: declarative.Alignment2D(walk.AlignHNearVNear),
			},
			declarative.LinkLabel{
				AssignTo:  &fb.FuncLabel,
				MinSize:   declarative.Size{Height: 50},
				MaxSize:   declarative.Size{Height: 50},
				Font:      declarative.Font{Family: "Roboto", PointSize: 9, Bold: true},
				Text:      instruction.FuncName,
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
			},
			declarative.Composite{
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
				MinSize:   declarative.Size{Height: 50},
				MaxSize:   declarative.Size{Height: 50},
				Layout:    declarative.HBox{MarginsZero: true, SpacingZero: true},
				Children: []declarative.Widget{
					declarative.Label{
						AssignTo:  &fb.NextIDLabel,
						Text:      "Next:",
						Font:      declarative.Font{Family: "Roboto", PointSize: 7},
						Alignment: declarative.Alignment2D(walk.AlignHFarVCenter),
					},
					declarative.NumberEdit{
						AssignTo:  &fb.NextID,
						MaxSize:   declarative.Size{Width: 25},
						Font:      declarative.Font{Family: "Roboto", PointSize: 7},
						Value:     float64(instruction.NextInstructionID),
						Decimals:  0,
						Alignment: declarative.Alignment2D(walk.AlignHNearVCenter),
						OnValueChanged: func() {
							instruction.NextInstructionID = int(fb.NextID.Value())
						},
					},
				},
			},
		},
	}).Create(declarative.NewBuilder(aw.ScrollView)); err != nil {
		return err
	}
	fb.Composite.SetXPixels(instruction.X)
	fb.Composite.SetYPixels(instruction.Y)
	fb.Composite.SetCursor(walk.CursorHand())

	// Disabling element position reset on other elements resizing
	fb.Composite.BoundsChanged().Attach(func() {
		// Checking if the position is automatically reseted
		if !pressed {
			fb.Composite.SetXPixels(instruction.X)
			fb.Composite.SetYPixels(instruction.Y)
		}
	})
	row++
	return nil
}

//DeleteFiberButton Delete the button from automate's screen & remove the associated instruction from the fiber
func (aw *Window) DeleteFiberButton(btn *button.Button) {
	dialogs.DeleteFiberButtonDialog.Run(aw.MainWindow)
	if dialogs.Dlg.Result() == 1 {
		for i := 0; i < len(NewButtons.Buttons); i++ {
			if btn == NewButtons.Buttons[i] {
				NewButtons.Buttons = append(NewButtons.Buttons[:i], NewButtons.Buttons[i+1:]...)
				currentFiber.Instructions = append(currentFiber.Instructions[:i], currentFiber.Instructions[i+1:]...)
				btn.NextID.Dispose()
				btn.NextIDLabel.Dispose()
				btn.FuncLabel.Dispose()
				btn.IDLabel.Dispose()
				btn.Composite.Dispose()
				return
			}
		}
		fmt.Println("ERROR: Unable to delete the fiber's instruction")
	}
}

//AddSavedFibersActions Add all the saved fibers to the My fibers's menu
func (aw *Window) AddSavedFibersActions() {
	root := "./saves"
	aw.Menu.FoldersMenu.Actions().Clear()

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			fullPath, _ := filepath.Abs(path)
			var file = filepath.Base(path)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			a := walk.NewAction()
			a.SetText(name)
			a.Triggered().Attach(func() { aw.InitOpenFiber(fullPath) })
			aw.Menu.FoldersMenu.Actions().Add(a)
		}
		return nil
	})
}

//NoFiberNameSet Show an error to tell user that the fiber as no name setted
func (aw *Window) NoFiberNameSet() {
	dialogs.NoFiberNameSetDialog.Run(aw.MainWindow)
	aw.FiberNameInput.SetFocus()
}

// InitCreateNewFiber Init a new fiber creation
func (aw *Window) InitCreateNewFiber() {
	if aw.IsFiberSaved() {
		aw.Clear()
		aw.CreateNewFiber()
	} else {
		dialogs.FiberNotSavedDialog.Run(aw.MainWindow)
		if dialogs.Dlg.Result() == 1 {
			aw.Clear()
			aw.CreateNewFiber()
		}
	}
}

// Clear Delete of the elements linked to the fiber
func (aw *Window) Clear() {
	aw.FiberNameInput.SetText("")
	aw.CleanScrollView()
	currentFiber.CleanFiber()
}

// CleanScrollView Dispose & recreate the MainWindow's ScrollView & empty the NewButtons slice
func (aw *Window) CleanScrollView() error {
	aw.ScrollView.Dispose()
	if err := (declarative.ScrollView{
		AssignTo:        &aw.ScrollView,
		Layout:          declarative.Grid{Columns: 1, Margins: declarative.Margins{Bottom: 2000, Right: 2000}},
		Background:      declarative.SolidColorBrush{Color: walk.RGB(11, 11, 11)},
		HorizontalFixed: false,
		VerticalFixed:   false,
	}).Create(declarative.NewBuilder(aw.MainWindow)); err != nil {
		return err
	}
	NewButtons.CleanButtons()
	return nil
}

// CreateNewFiber Create a new fiber
func (aw *Window) CreateNewFiber() {
	data, dialog := packages.CreateNewDialog("Flow", "Start")

	newInstruction := &fiber.Instruction{
		ID:                0,
		Package:           "Flow",
		FuncName:          "Start",
		X:                 0,
		Y:                 0,
		NextInstructionID: 1,
		InstructionData:   data,
	}
	currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
	aw.CreateFiberButton(newInstruction, dialog, false)
}

//InitSaveFiber Check if the fiber can be saved & save it if possible
func (aw *Window) InitSaveFiber() error {
	name := aw.FiberNameInput.Text()
	if name == "" {
		aw.NoFiberNameSet()
		return nil
	}

	// Checking if the fiber is already saved
	if !aw.IsFiberSaved() {
		aw.SaveFiber("saves/" + name)
		aw.AddSavedFibersActions()
	}
	return nil
}

//ExportFiber Open a dialog window for user directory selection
func (aw *Window) ExportFiber() error {
	name := aw.FiberNameInput.Text()
	if name == "" {
		aw.NoFiberNameSet()
		return nil
	}
	dlg := new(walk.FileDialog)

	dlg.FilePath = dlg.FilePath + name

	dlg.Filter = "JSON File (*.json)|*.json"
	dlg.Title = "Export a fiber"

	if ok, err := dlg.ShowSave(aw.MainWindow); err != nil {
		return err
	} else if !ok {
		return nil
	}

	aw.SaveFiber(dlg.FilePath)
	return nil
}

//SaveFiber save the current fiber to the path parameter
func (aw *Window) SaveFiber(path string) {
	fullPath := path + ".json"
	file, _ := json.Marshal(currentFiber)
	ioutil.WriteFile(fullPath, file, 0644)
}

// InitImportFiber prepare for the importation of a fiber
func (aw *Window) InitImportFiber() {
	if aw.IsFiberSaved() {
		aw.ImportFiber()
	} else {
		dialogs.FiberNotSavedDialog.Run(aw.MainWindow)
		if dialogs.Dlg.Result() == 1 {
			aw.ImportFiber()
		}
	}
}

//ImportFiber Open a dialog window for user file selection
func (aw *Window) ImportFiber() error {
	dlg := new(walk.FileDialog)

	dlg.Filter = "JSON File (*.json)|*.json"
	dlg.Title = "Open a fiber"

	if ok, err := dlg.ShowOpen(aw.MainWindow); err != nil {
		return err
	} else if !ok {
		return nil
	}

	aw.OpenFiber(dlg.FilePath)

	return nil
}

// InitOpenFiber prepare for the opening of a fiber
func (aw *Window) InitOpenFiber(path string) {
	if aw.IsFiberSaved() {
		aw.OpenFiber(path)
	} else {
		dialogs.FiberNotSavedDialog.Run(aw.MainWindow)
		if dialogs.Dlg.Result() == 1 {
			aw.OpenFiber(path)
		}
	}
}

//OpenFiber Open a saved fiber from the menu My Fibers
func (aw *Window) OpenFiber(path string) {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println("ERROR: Unable to open the saved fiber")
		return
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var loadingFiber fiber.LoadingFiber
	err = json.Unmarshal(byteValue, &loadingFiber)

	aw.Clear()

	aw.FiberNameInput.SetText(loadingFiber.Name)

	currentFiber.Name = loadingFiber.Name
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	for _, instruction := range loadingFiber.Instructions {
		structure := packages.PackageDecode(instruction)

		err := json.Unmarshal(instruction.InstructionData, structure)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		newInstruction := &fiber.Instruction{
			ID:                instruction.ID,
			Package:           instruction.Package,
			FuncName:          instruction.FuncName,
			X:                 instruction.X,
			Y:                 instruction.Y,
			NextInstructionID: instruction.NextInstructionID,
			InstructionData:   structure,
		}
		_, dialog := packages.CreateNewDialog(instruction.Package, instruction.FuncName, structure)

		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog, true)
	}
}

// IsFiberSaved Check if the fiber is saved
func (aw *Window) IsFiberSaved() bool {
	exist := false
	name := aw.FiberNameInput.Text()
	currentFiber.Name = name
	file, _ := json.Marshal(currentFiber)

	root := "./saves"

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			fullPath, _ := filepath.Abs(path)
			jsonFile, _ := os.Open(fullPath)
			byteValue, _ := ioutil.ReadAll(jsonFile)
			if bytes.Compare(file, byteValue) == 0 {
				exist = true
			}
		}
		return nil
	})

	return exist
}
