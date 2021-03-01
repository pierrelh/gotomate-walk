package fiber

import (
	"encoding/json"
	"fmt"
	"gotomate/packages/battery"
	"gotomate/packages/clipboard"
	"gotomate/packages/keyboard"
	"gotomate/packages/log"
	"gotomate/packages/mouse"
	"gotomate/packages/notification"
	"gotomate/packages/screen"
	"gotomate/packages/sleep"
	"gotomate/packages/systime"
	"reflect"
	"regexp"
	"strings"

	batType "github.com/distatus/battery"
	"gopkg.in/toast.v1"
)

//NewFiber Define the new automate's fiber
var NewFiber = new(Fiber)
var running = 0
var finished = make(chan bool)
var stop = make(chan bool)

// LoadingFiber Initialize the loading fiber structure
type LoadingFiber struct {
	Name         string
	Instructions []*LoadingInstruction
}

// LoadingInstruction Initialize a loading fiber's instruction
type LoadingInstruction struct {
	Package  string
	FuncName string
	Data     json.RawMessage
}

// Fiber Initialize the fiber structure
type Fiber struct {
	Name         string
	Instructions []*Instruction
}

// Instruction Initialize a fiber's instruction
type Instruction struct {
	Package  string
	FuncName string
	Data     interface{}
}

//CleanFiber Delete all the instructions of the current fiber
func (fiber *Fiber) CleanFiber() {
	p := reflect.ValueOf(fiber).Elem()
	p.Set(reflect.Zero(p.Type()))
}

//StopFiber stop the current fiber
func (fiber *Fiber) StopFiber() {
	finished <- true
	stop <- true
}

//RunFiber run the current fiber
func (fiber *Fiber) RunFiber() {
	running++
	if running > 1 {
		fmt.Println("FIBER: A fiber is already running")
	} else {
		fmt.Println("| Fiber Start |")

		for i := 0; i < len(fiber.Instructions); i++ {
			select {
			case <-stop:
				fmt.Println("| Fiber Stopped |")
				running = 0
				return
			default:
				instruction := fiber.Instructions[i]
				data := reflect.ValueOf(instruction.Data).Elem()
				switch instruction.Package {
				case "Sleep":
					switch instruction.FuncName {
					case "Sleep":
						duration := data.FieldByName("Duration").Interface().(float64)
						go sleep.Sleep(duration, finished)
						<-finished
					case "MilliSleep":
						duration := data.FieldByName("Duration").Interface().(float64)
						go sleep.MilliSleep(duration, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Mouse":
					switch instruction.FuncName {
					case "Click":
						button := data.FieldByName("MouseButtonName").Interface().(string)
						go mouse.Click(button, finished)
						<-finished
					case "Scroll":
						x := data.FieldByName("X").Interface().(int)
						y := data.FieldByName("Y").Interface().(int)
						go mouse.Scroll(x, y, finished)
						<-finished
					case "Move":
						x := data.FieldByName("X").Interface().(int)
						y := data.FieldByName("Y").Interface().(int)
						go mouse.Move(x, y, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Keyboard":
					switch instruction.FuncName {
					case "Tap":
						input := data.FieldByName("Input").Interface().(string)
						special := []string{}
						if err := data.FieldByName("Special1").Interface().(string); err != "" {
							special = append(special, err)
						}
						if err := data.FieldByName("Special2").Interface().(string); err != "" {
							special = append(special, err)
						}
						go keyboard.Tap(input, special, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Clipboard":
					switch instruction.FuncName {
					case "Write":
						content := data.FieldByName("Content").Interface().(string)
						go clipboard.Write(content, finished)
						<-finished
					case "Read":
						go func() {
							content, _ := clipboard.Read(finished)
							data.FieldByName("Value").SetString(content)
						}()
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Log":
					switch instruction.FuncName {
					case "Print":
						logValue := data.FieldByName("Log").Interface().(string)
						matched, _ := regexp.MatchString(`=.*`, logValue)
						msg := data.FieldByName("Log").Interface()
						if matched {
							varSlice := strings.Split(logValue, "=")
							for i := 0; i < len(varSlice); i++ {
								if varSlice[i] != "" {
									varInstruction := fiber.VarSearch(regexp.QuoteMeta(varSlice[i]))
									if varInstruction != nil {
										val := reflect.ValueOf(varInstruction.Data).Elem()
										msg = val.FieldByName("Value").Interface()
									} else {
										fmt.Println("ERROR: Unable to find the fiber's var: ", regexp.QuoteMeta(varSlice[i]))
										msg = ""
									}
								}
							}
						}
						go log.Print(msg, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Notification":
					switch instruction.FuncName {
					case "Create":
						title := data.FieldByName("Title").Interface().(string)
						msg := data.FieldByName("Message").Interface().(string)
						actions := data.FieldByName("Actions").Interface().([]toast.Action)
						go notification.Create(title, msg, actions, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Battery":
					switch instruction.FuncName {
					case "GetBattery":
						go func() {
							bat := battery.GetBattery(finished)
							data.FieldByName("Value").Set(reflect.ValueOf(bat))
						}()
						<-finished
					case "GetBatteryState":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								state := battery.GetBatteryState(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(state))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryPercentage":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								percentage := battery.GetBatteryPercentage(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(percentage))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryRemainingTime":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								time := battery.GetBatteryRemainingTime(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(time))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryChargeRate":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								rate := battery.GetBatteryChargeRate(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(rate))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryCurrentCapacity":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								capacity := battery.GetBatteryCurrentCapacity(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryLastFullCapacity":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								capacity := battery.GetBatteryLastFullCapacity(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryDesignCapacity":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								capacity := battery.GetBatteryDesignCapacity(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryVoltage":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								voltage := battery.GetBatteryVoltage(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(voltage))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryDesignVoltage":
						go func() {
							batName := data.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*batType.Battery)
								voltage := battery.GetBatteryDesignVoltage(bat, finished)
								data.FieldByName("Value").Set(reflect.ValueOf(voltage))
							} else {
								data.FieldByName("Value").Set(reflect.Zero(data.FieldByName("Value").Type()))
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Systime":
					switch instruction.FuncName {
					case "GetCurrentSysClock":
						go func() {
							h, m, s := systime.GetCurrentSysClock(finished)
							clock := [3]int{h, m, s}
							data.FieldByName("Value").Set(reflect.ValueOf(clock))
						}()
						<-finished
					case "GetCurrentSysTime":
						go func() {
							time := systime.GetCurrentSysTime(finished)
							data.FieldByName("Value").Set(reflect.ValueOf(time))
						}()
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Screen":
					switch instruction.FuncName {
					case "GetMouseColor":
						go func() {
							value := screen.GetMouseColor(finished)
							data.FieldByName("Value").Set(reflect.ValueOf(value))
						}()
						<-finished
					case "GetPixelColor":
						go func() {
							x := data.FieldByName("X").Interface().(int)
							y := data.FieldByName("Y").Interface().(int)
							value := screen.GetPixelColor(x, y, finished)
							data.FieldByName("Value").Set(reflect.ValueOf(value))
						}()
						<-finished
					case "SaveCapture":
						path := data.FieldByName("Path").Interface().(string)
						go screen.SaveCapture(finished, path)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				default:
					fmt.Println("FIBER: This package is not integrated yet: " + instruction.Package)
					continue
				}
			}
		}
		fmt.Println("| Fiber Finished |")
		running = 0
	}
}
