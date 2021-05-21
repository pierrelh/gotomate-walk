package fiber

import (
	"encoding/json"
	"fmt"

	// DON'T REMOVE ME / New packages inserted here
	algorithmic "gotomate/fiber/packages/Algorithmic"
	arithmetic "gotomate/fiber/packages/Arithmetic"
	array "gotomate/fiber/packages/Array"
	battery "gotomate/fiber/packages/Battery"
	clipboard "gotomate/fiber/packages/Clipboard"
	conversion "gotomate/fiber/packages/Conversion"
	define "gotomate/fiber/packages/Define"
	file "gotomate/fiber/packages/File"
	flow "gotomate/fiber/packages/Flow"
	input "gotomate/fiber/packages/Input"
	keyboard "gotomate/fiber/packages/Keyboard"
	log "gotomate/fiber/packages/Log"
	mouse "gotomate/fiber/packages/Mouse"
	notification "gotomate/fiber/packages/Notification"
	process "gotomate/fiber/packages/Process"
	scraping "gotomate/fiber/packages/Scraping"
	screen "gotomate/fiber/packages/Screen"
	sleep "gotomate/fiber/packages/Sleep"
	sound "gotomate/fiber/packages/Sound"
	systime "gotomate/fiber/packages/Systime"
	"gotomate/fiber/variable"
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
	if running == 1 {
		fmt.Println("| Fiber Stopped |")
		finished <- true
		running = 0
		stop <- true
	} else {
		fmt.Println("FIBER WARNING: No running fiber")
	}
}

//RunFiber run the current fiber
func (fiber *Fiber) RunFiber() {
	running++
	if running > 1 {
		fmt.Println("FIBER WARNING: A fiber is already running")
	} else {
		instruction := fiber.Instructions[0]
		variable.FiberVariable = nil

		for {
			select {
			case <-stop:
				return
			default:
				nextID := -1
				funcName := instruction.FuncName
				instructionData := reflect.ValueOf(instruction.InstructionData).Elem()
				switch instruction.Package {
				case "Flow":
					ended := flow.Processing(funcName, instructionData, finished)
					if ended {
						running = 0
						return
					}
				// DON'T REMOVE ME / New processing inserted here
				case "Algorithmic":
					nextID = algorithmic.Processing(funcName, instructionData, finished)
				case "Arithmetic":
					nextID = arithmetic.Processing(funcName, instructionData, finished)
				case "Array":
					nextID = array.Processing(funcName, instructionData, finished)
				case "Battery":
					nextID = battery.Processing(funcName, instructionData, finished)
				case "Clipboard":
					nextID = clipboard.Processing(funcName, instructionData, finished)
				case "Conversion":
					nextID = conversion.Processing(funcName, instructionData, finished)
				case "Define":
					nextID = define.Processing(funcName, instructionData, finished)
				case "File":
					nextID = file.Processing(funcName, instructionData, finished)
				case "Input":
					nextID = input.Processing(funcName, instructionData, finished)
				case "Keyboard":
					nextID = keyboard.Processing(funcName, instructionData, finished)
				case "Log":
					nextID = log.Processing(funcName, instructionData, finished)
				case "Mouse":
					nextID = mouse.Processing(funcName, instructionData, finished)
				case "Notification":
					nextID = notification.Processing(funcName, instructionData, finished)
				case "Process":
					nextID = process.Processing(funcName, instructionData, finished)
				case "Scraping":
					nextID = scraping.Processing(funcName, instructionData, finished)
				case "Screen":
					nextID = screen.Processing(funcName, instructionData, finished)
				case "Sleep":
					nextID = sleep.Processing(funcName, instructionData, finished)
				case "Sound":
					nextID = sound.Processing(funcName, instructionData, finished)
				case "Systime":
					nextID = systime.Processing(funcName, instructionData, finished)
				default:
					fmt.Println("FIBER WARNING: This package is not integrated yet: " + instruction.Package)
					continue
				}

				if nextID == -1 {
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
				} else {
					instruction = fiber.Instructions[nextID]
				}
			}
		}
	}
}
