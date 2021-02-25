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
	"gotomate/packages/sleep"
	"gotomate/packages/systime"
	"reflect"
	"regexp"
	"strings"

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
				switch instruction.Package {
				case "Sleep":
					switch instruction.FuncName {
					case "Sleep":
						val := reflect.ValueOf(instruction.Data).Elem()
						duration := val.FieldByName("Duration").Interface().(float64)
						go sleep.Sleep(duration, finished)
						<-finished
					case "MilliSleep":
						val := reflect.ValueOf(instruction.Data).Elem()
						duration := val.FieldByName("Duration").Interface().(float64)
						go sleep.MilliSleep(duration, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Mouse":
					switch instruction.FuncName {
					case "Click":
						val := reflect.ValueOf(instruction.Data).Elem()
						button := val.FieldByName("MouseButtonName").Interface().(string)
						go mouse.Click(button, finished)
						<-finished
					case "Scroll":
						val := reflect.ValueOf(instruction.Data).Elem()
						x := val.FieldByName("X").Interface().(int)
						y := val.FieldByName("Y").Interface().(int)
						go mouse.Scroll(x, y, finished)
						<-finished
					case "Move":
						val := reflect.ValueOf(instruction.Data).Elem()
						x := val.FieldByName("X").Interface().(int)
						y := val.FieldByName("Y").Interface().(int)
						go mouse.Move(x, y, finished)
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Keyboard":
					switch instruction.FuncName {
					case "Tap":
						val := reflect.ValueOf(instruction.Data).Elem()
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
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Clipboard":
					switch instruction.FuncName {
					case "Write":
						val := reflect.ValueOf(instruction.Data).Elem()
						content := val.FieldByName("Content").Interface().(string)
						go clipboard.Write(content, finished)
						<-finished
					case "Read":
						go func() {
							content, _ := clipboard.Read(finished)
							val := reflect.ValueOf(instruction.Data).Elem()
							val.FieldByName("Value").SetString(content)
						}()
						<-finished
					default:
						fmt.Println("FIBER: This function is not integrated yet: " + instruction.FuncName)
						continue
					}
				case "Log":
					switch instruction.FuncName {
					case "Print":
						val := reflect.ValueOf(instruction.Data).Elem()
						logValue := val.FieldByName("Log").Interface().(string)
						matched, _ := regexp.MatchString(`=.*`, logValue)
						msg := val.FieldByName("Log").Interface()
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
						val := reflect.ValueOf(instruction.Data).Elem()
						title := val.FieldByName("Title").Interface().(string)
						msg := val.FieldByName("Message").Interface().(string)
						actions := val.FieldByName("Actions").Interface().([]toast.Action)
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
							val := reflect.ValueOf(instruction.Data).Elem()
							val.FieldByName("Value").Set(reflect.ValueOf(bat))
						}()
						<-finished
					case "GetBatteryState":
						go func() {
							stateInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := stateInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								state := battery.GetBatteryState(bat, finished)
								stateInstruction.FieldByName("Value").Set(reflect.ValueOf(state))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryPercentage":
						go func() {
							percentageInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := percentageInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								percentage := battery.GetBatteryPercentage(bat, finished)
								percentageInstruction.FieldByName("Value").Set(reflect.ValueOf(percentage))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryRemainingTime":
						go func() {
							remainingTimeInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := remainingTimeInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								time := battery.GetBatteryRemainingTime(bat, finished)
								remainingTimeInstruction.FieldByName("Value").Set(reflect.ValueOf(time))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryChargeRate":
						go func() {
							chargeRateInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := chargeRateInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								rate := battery.GetBatteryChargeRate(bat, finished)
								chargeRateInstruction.FieldByName("Value").Set(reflect.ValueOf(rate))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryCurrentCapacity":
						go func() {
							currentCapacityInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := currentCapacityInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								capacity := battery.GetBatteryCurrentCapacity(bat, finished)
								currentCapacityInstruction.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryLastFullCapacity":
						go func() {
							lastCapacityInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := lastCapacityInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								capacity := battery.GetBatteryLastFullCapacity(bat, finished)
								lastCapacityInstruction.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryDesignCapacity":
						go func() {
							designCapacityInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := designCapacityInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								capacity := battery.GetBatteryDesignCapacity(bat, finished)
								designCapacityInstruction.FieldByName("Value").Set(reflect.ValueOf(capacity))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryVoltage":
						go func() {
							voltageInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := voltageInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								voltage := battery.GetBatteryVoltage(bat, finished)
								voltageInstruction.FieldByName("Value").Set(reflect.ValueOf(voltage))
							} else {
								fmt.Println("FIBER: Unable to find the battery named: ", batName)
								finished <- true
							}
						}()
						<-finished
					case "GetBatteryDesignVoltage":
						go func() {
							designVoltageInstruction := reflect.ValueOf(instruction.Data).Elem()
							batName := designVoltageInstruction.FieldByName("BatteryName").Interface().(string)
							batInstruction := fiber.VarSearch(batName)
							if batInstruction != nil {
								val := reflect.ValueOf(batInstruction.Data).Elem()
								bat := val.FieldByName("Value").Interface().(*battery.Battery)
								voltage := battery.GetBatteryDesignVoltage(bat, finished)
								designVoltageInstruction.FieldByName("Value").Set(reflect.ValueOf(voltage))
							} else {
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
							val := reflect.ValueOf(instruction.Data).Elem()
							val.FieldByName("Value").Set(reflect.ValueOf(clock))
						}()
						<-finished
					case "GetCurrentSysTime":
						go func() {
							time := systime.GetCurrentSysTime(finished)
							val := reflect.ValueOf(instruction.Data).Elem()
							val.FieldByName("Value").Set(reflect.ValueOf(time))
						}()
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
