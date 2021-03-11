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

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//NewButtons Create the automate's list of buttons
var NewButtons = button.NewButtons
var currentFiber = fiber.NewFiber
var pressed = false

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

		data, dialog := packages.CreateNewDialog(packageName, funcName)

		newInstruction := &fiber.Instruction{
			Package:         packageName,
			FuncName:        funcName,
			InstructionData: data,
		}
		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog)
	}
}

// CreateFiberButton Create a new Fiberbutton in the fiber
func (aw *Window) CreateFiberButton(instruction *fiber.Instruction, dialog *packages.Dialog) error {
	fb := new(button.Button)
	NewButtons.Buttons = append(NewButtons.Buttons, fb)
	fb.Dialog = dialog
	bmp, err := walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/fiber/packages/" + instruction.Package + "/icon.png")
	if err != nil {
		bmp, _ = walk.NewBitmapFromFile(walk.Resources.RootDirPath() + "/img/default.png")
	}

	if err := (declarative.Composite{
		AssignTo:   &fb.Composite,
		Layout:     declarative.VBox{MarginsZero: true, SpacingZero: true},
		Background: declarative.BitmapBrush{Image: bmp},
		Alignment:  declarative.Alignment2D(walk.AlignHCenterVCenter),
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
					aw.DeleteFiberButton(fb)
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

//DeleteFiberButton Delete the button from automate's screen & remove the associated instruction from the fiber
func (aw *Window) DeleteFiberButton(btn *button.Button) {
	dialogs.DeleteFiberButtonDialog.Run(aw.MainWindow)
	if dialogs.Dlg.Result() == 1 {
		for i := 0; i < len(NewButtons.Buttons); i++ {
			if btn == NewButtons.Buttons[i] {
				NewButtons.Buttons = append(NewButtons.Buttons[:i], NewButtons.Buttons[i+1:]...)
				currentFiber.Instructions = append(currentFiber.Instructions[:i], currentFiber.Instructions[i+1:]...)
				btn.Composite.Dispose()
				btn.LinkLabel.Dispose()
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

// CreateNewFiber Init a new fiber creation
func (aw *Window) CreateNewFiber() {
	if aw.IsFiberSaved() {
		aw.FiberNameInput.SetText("")
		NewButtons.CleanButtons()
		currentFiber.CleanFiber()
	} else {
		dialogs.FiberNotSavedDialog.Run(aw.MainWindow)
		if dialogs.Dlg.Result() == 1 {
			aw.FiberNameInput.SetText("")
			NewButtons.CleanButtons()
			currentFiber.CleanFiber()
		}
	}
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
	aw.FiberNameInput.SetText(loadingFiber.Name)

	NewButtons.CleanButtons()
	currentFiber.CleanFiber()

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
			Package:         instruction.Package,
			FuncName:        instruction.FuncName,
			InstructionData: structure,
		}
		_, dialog := packages.CreateNewDialog(instruction.Package, instruction.FuncName, structure)

		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog)
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
