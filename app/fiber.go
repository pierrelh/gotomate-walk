package app

import (
	"fmt"
	"gotomate/battery"
	"gotomate/clipboard"
	"gotomate/keyboard"
	"gotomate/log"
	"gotomate/mouse"
	"gotomate/notification"
	"gotomate/sleep"
	"reflect"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"gopkg.in/toast.v1"
)

// Fiber Initialize the fiber structure
type Fiber struct {
	Instructions []*FiberInstruction
}

// FiberInstruction Initialize a fiber's instruction
type FiberInstruction struct {
	Package  string
	FuncName string
	Button   *FiberButton
}

// FiberButton Setting fiber's buttons structure
type FiberButton struct {
	*walk.Composite
	*walk.ImageView
	*walk.LinkLabel
	DialogWindow declarative.Dialog
}

// FiberComposite Setting FiberButton's composite & it's dialog
type FiberComposite struct {
	*walk.Composite
	DialogWindow declarative.Dialog
}

// FiberImageView Setting FiberButton's imageview & it's dialog
type FiberImageView struct {
	*walk.ImageView
	DialogWindow declarative.Dialog
}

// FiberLinkLabel Setting FiberButton's linklabel & it's dialog
type FiberLinkLabel struct {
	*walk.LinkLabel
	DialogWindow declarative.Dialog
}

//WndProc setting the window event of the composite element
func (fb *FiberComposite) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.Composite.WndProc(hwnd, msg, wParam, lParam)
}

//WndProc setting the window event of the imageview element
func (fb *FiberImageView) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.ImageView.WndProc(hwnd, msg, wParam, lParam)
}

//WndProc setting the window event of the linklabel element
func (fb *FiberLinkLabel) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		go fb.DialogWindow.Run(aw.mw)
	}

	return fb.LinkLabel.WndProc(hwnd, msg, wParam, lParam)
}

func runFiber() {
	fmt.Println("| Fiber Start |")

	for i := 0; i < len(fiber.Instructions); i++ {
		finished := make(chan bool)
		instruction := fiber.Instructions[i]
		switch instruction.Package {
		case "Sleep":
			switch instruction.FuncName {
			case "Sleep":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				duration := val.FieldByName("Duration").Interface().(float64)
				go sleep.Sleep(duration, finished)
				<-finished
			case "MilliSleep":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				duration := val.FieldByName("Duration").Interface().(float64)
				go sleep.MilliSleep(duration, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Mouse":
			switch instruction.FuncName {
			case "Click":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				button := val.FieldByName("MouseButtonName").Interface().(string)
				go mouse.Click(button, finished)
				<-finished
			case "Scroll":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				x := val.FieldByName("X").Interface().(int)
				y := val.FieldByName("Y").Interface().(int)
				go mouse.Scroll(x, y, finished)
				<-finished
			case "Move":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				x := val.FieldByName("X").Interface().(int)
				y := val.FieldByName("Y").Interface().(int)
				go mouse.Move(x, y, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Keyboard":
			switch instruction.FuncName {
			case "Tap":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				input := val.FieldByName("Input").Interface().(string)
				special := []string{}
				if err := val.FieldByName("Special1").Interface().(string); err != "" {
					special = append(special, err)
				}
				if err := val.FieldByName("Special2").Interface().(string); err != "" {
					special = append(special, err)
				}
				go keyboard.Tap(input, special, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Clipboard":
			switch instruction.FuncName {
			case "Write":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				content := val.FieldByName("Content").Interface().(string)
				go clipboard.Write(content, finished)
				<-finished
			case "Read":
				go func() {
					content, _ := clipboard.Read(finished)
					val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					val.FieldByName("Content").SetString(content)
				}()
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Log":
			switch instruction.FuncName {
			case "Print":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				msg := val.FieldByName("Log").Interface().(string)
				go log.Print(msg, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Notification":
			switch instruction.FuncName {
			case "Create":
				val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
				title := val.FieldByName("Title").Interface().(string)
				msg := val.FieldByName("Message").Interface().(string)
				actions := val.FieldByName("Actions").Interface().([]toast.Action)
				go notification.Create(title, msg, actions, finished)
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		case "Battery":
			switch instruction.FuncName {
			case "GetBattery":
				go func() {
					bat := battery.GetBattery(finished)
					val := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					val.FieldByName("Battery").Set(reflect.ValueOf(bat))
				}()
				<-finished
			case "GetBatteryState":
				go func() {
					stateInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := stateInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						state := battery.GetBatteryState(bat, finished)
						stateInstruction.FieldByName("State").Set(reflect.ValueOf(state))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryPercentage":
				go func() {
					percentageInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := percentageInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						percentage := battery.GetBatteryPercentage(bat, finished)
						percentageInstruction.FieldByName("Percentage").Set(reflect.ValueOf(percentage))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryRemainingTime":
				go func() {
					remainingTimeInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := remainingTimeInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						time := battery.GetBatteryRemainingTime(bat, finished)
						remainingTimeInstruction.FieldByName("RemainingTime").Set(reflect.ValueOf(time))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryChargeRate":
				go func() {
					chargeRateInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := chargeRateInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						rate := battery.GetBatteryChargeRate(bat, finished)
						chargeRateInstruction.FieldByName("ChargeRate").Set(reflect.ValueOf(rate))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryCurrentCapacity":
				go func() {
					currentCapacityInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := currentCapacityInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						capacity := battery.GetBatteryCurrentCapacity(bat, finished)
						currentCapacityInstruction.FieldByName("CurrentCapacity").Set(reflect.ValueOf(capacity))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryLastFullCapacity":
				go func() {
					lastCapacityInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := lastCapacityInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						capacity := battery.GetBatteryLastFullCapacity(bat, finished)
						lastCapacityInstruction.FieldByName("LastFullCapacity").Set(reflect.ValueOf(capacity))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryDesignCapacity":
				go func() {
					designCapacityInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := designCapacityInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						capacity := battery.GetBatteryDesignCapacity(bat, finished)
						designCapacityInstruction.FieldByName("DesignCapacity").Set(reflect.ValueOf(capacity))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryVoltage":
				go func() {
					voltageInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := voltageInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						voltage := battery.GetBatteryVoltage(bat, finished)
						voltageInstruction.FieldByName("Voltage").Set(reflect.ValueOf(voltage))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			case "GetBatteryDesignVoltage":
				go func() {
					designVoltageInstruction := reflect.ValueOf(instruction.Button.DialogWindow.DataBinder.DataSource).Elem()
					batName := designVoltageInstruction.FieldByName("BatteryName").Interface().(string)
					batInstruction := varSearch(batName)
					if batInstruction != nil {
						val := reflect.ValueOf(batInstruction.Button.DialogWindow.DataBinder.DataSource).Elem()
						bat := val.FieldByName("Battery").Interface().(*battery.Battery)
						voltage := battery.GetBatteryDesignVoltage(bat, finished)
						designVoltageInstruction.FieldByName("DesignVoltage").Set(reflect.ValueOf(voltage))
					} else {
						fmt.Println("Unable to find the battery named: ", batName)
						finished <- true
					}
				}()
				<-finished
			default:
				fmt.Println("This function is not integrated yet: " + instruction.FuncName)
				continue
			}
		default:
			fmt.Println("This package is not integrated yet: " + instruction.Package)
			continue
		}
	}
	fmt.Println("| Fiber Finished |")
}
