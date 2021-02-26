package automate

import (
	"encoding/json"
	"fmt"
	"gotomate/app/fiber"
	"gotomate/packages"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

//AddSavedFibersActions Add all the saved fibers to the My fibers's menu
func (aw *Window) AddSavedFibersActions(currentFiber *fiber.Fiber) {
	root := "./saves"
	aw.Menu.Folders.Actions().Clear()

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			var file = filepath.Base(path)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			a := walk.NewAction()
			a.SetText(name)
			a.Triggered().Attach(func() { aw.OpenSavedFiber(path, currentFiber) })
			aw.Menu.Folders.Actions().Add(a)
		}
		return nil
	})
}

//OpenSavedFiber Open a saved fiber from the menu My Fibers
func (aw *Window) OpenSavedFiber(path string, currentFiber *fiber.Fiber) {
	jsonFile, err := os.Open("./" + path)

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
		var structure interface{}
		switch instruction.FuncName {
		case "Log":
			structure = new(packages.LogPrint)
		case "Sleep":
			structure = new(packages.Sleep)
		case "MilliSleep":
			structure = new(packages.MilliSleep)
		case "Click":
			structure = new(packages.MouseClick)
		case "Scroll":
			structure = new(packages.MouseScroll)
		case "Move":
			structure = new(packages.MouseMove)
		case "Tap":
			structure = new(packages.KeyboardTap)
		case "Write":
			structure = new(packages.ClipboardWrite)
		case "Read":
			structure = new(packages.ClipboardRead)
		case "Print":
			structure = new(packages.LogPrint)
		case "Create":
			structure = new(packages.NotificationCreate)
		case "GetBattery":
			structure = new(packages.UserBattery)
		case "GetBatteryState":
			structure = new(packages.BatteryState)
		case "GetBatteryPercentage":
			structure = new(packages.BatteryPercentage)
		case "GetBatteryRemainingTime":
			structure = new(packages.BatteryRemainingTime)
		case "GetBatteryChargeRate":
			structure = new(packages.BatteryChargeRate)
		case "GetBatteryCurrentCapacity":
			structure = new(packages.BatteryCurrentCapacity)
		case "GetBatteryLastFullCapacity":
			structure = new(packages.BatteryLastFullCapacity)
		case "GetBatteryDesignCapacity":
			structure = new(packages.BatteryDesignCapacity)
		case "GetBatteryVoltage":
			structure = new(packages.BatteryVoltage)
		case "GetBatteryDesignVoltage":
			structure = new(packages.BatteryDesignVoltage)
		case "GetCurrentSysClock":
			structure = new(packages.SysClock)
		case "GetCurrentSysTime":
			structure = new(packages.SysTime)
		default:
			fmt.Println("ERROR: Unable to find the function")
			structure = nil
		}

		err := json.Unmarshal(instruction.Data, structure)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		newInstruction := &fiber.Instruction{
			Package:  instruction.Package,
			FuncName: instruction.FuncName,
			Data:     structure,
		}
		_, dialog := packages.CreateNewDialog(instruction.FuncName, structure)

		currentFiber.Instructions = append(currentFiber.Instructions, newInstruction)
		aw.CreateFiberButton(newInstruction, dialog)
	}
}

//SaveFiber save the current fiber
func (aw *Window) SaveFiber(currentFiber *fiber.Fiber) {
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
		currentFiber.Name = name
		path := "saves/" + name + ".json"
		file, _ := json.Marshal(currentFiber)
		ioutil.WriteFile(path, file, 0644)
		aw.AddSavedFibersActions(currentFiber)
	}
}
