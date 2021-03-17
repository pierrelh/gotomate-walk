package fiber

import (
	"encoding/json"
	"fmt"
	battery "gotomate/fiber/packages/Battery"
	clipboard "gotomate/fiber/packages/Clipboard"
	flow "gotomate/fiber/packages/Flow"
	input "gotomate/fiber/packages/Input"
	keyboard "gotomate/fiber/packages/Keyboard"
	log "gotomate/fiber/packages/Log"
	mouse "gotomate/fiber/packages/Mouse"
	notification "gotomate/fiber/packages/Notification"
	process "gotomate/fiber/packages/Process"
	screen "gotomate/fiber/packages/Screen"
	sleep "gotomate/fiber/packages/Sleep"
	systime "gotomate/fiber/packages/Systime"
	"gotomate/fiber/value"
	"reflect"
	"sort"
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
	ID                int
	Package           string
	FuncName          string
	X                 int
	Y                 int
	NextInstructionID int
	InstructionData   json.RawMessage
}

// Fiber Initialize the fiber structure
type Fiber struct {
	Name         string
	Instructions []*Instruction
}

// Instruction Initialize a fiber's instruction
type Instruction struct {
	ID                int
	Package           string
	FuncName          string
	X                 int
	Y                 int
	NextInstructionID int
	InstructionData   interface{}
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
		fmt.Println("FIBER WARNING: A fiber is already running")
	} else {
		instruction := fiber.Instructions[0]
		value.FiberValues = nil

		for {
			select {
			case <-stop:
				fmt.Println("| Fiber Stopped |")
				running = 0
				return
			default:
				funcName := instruction.FuncName
				instructionData := reflect.ValueOf(instruction.InstructionData).Elem()
				switch instruction.Package {
				case "Flow":
					ended := flow.Processing(funcName, instructionData, finished)
					if ended {
						running = 0
						return
					}
				case "Sleep":
					sleep.Processing(funcName, instructionData, finished)
				case "Mouse":
					mouse.Processing(funcName, instructionData, finished)
				case "Keyboard":
					keyboard.Processing(funcName, instructionData, finished)
				case "Clipboard":
					clipboard.Processing(funcName, instructionData, finished)
				case "Log":
					log.Processing(funcName, instructionData, finished)
				case "Notification":
					notification.Processing(funcName, instructionData, finished)
				case "Battery":
					battery.Processing(funcName, instructionData, finished)
				case "Systime":
					systime.Processing(funcName, instructionData, finished)
				case "Screen":
					screen.Processing(funcName, instructionData, finished)
				case "Process":
					process.Processing(funcName, instructionData, finished)
				case "Input":
					input.Processing(funcName, instructionData, finished)
				default:
					fmt.Println("FIBER WARNING: This package is not integrated yet: " + instruction.Package)
					continue
				}

				idx := sort.Search(len(fiber.Instructions), func(i int) bool {
					return fiber.Instructions[i].ID >= instruction.NextInstructionID
				})
				if idx == len(fiber.Instructions) {
					fmt.Println("FIBER FATAL ERROR: The instruction with the id", instruction.NextInstructionID, "has no been founded")
					fmt.Println("| Fiber Finished at Fatal Error |")
					running = 0
					return
				} else {
					instruction = fiber.Instructions[idx]
				}
			}
		}
	}
}
