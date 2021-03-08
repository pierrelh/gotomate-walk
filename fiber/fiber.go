package fiber

import (
	"encoding/json"
	"fmt"
	"gotomate/fiber/packages/battery"
	"gotomate/fiber/packages/clipboard"
	"gotomate/fiber/packages/input"
	"gotomate/fiber/packages/keyboard"
	"gotomate/fiber/packages/log"
	"gotomate/fiber/packages/mouse"
	"gotomate/fiber/packages/notification"
	"gotomate/fiber/packages/process"
	"gotomate/fiber/packages/screen"
	"gotomate/fiber/packages/sleep"
	"gotomate/fiber/packages/systime"
	"reflect"
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
	Package         string
	FuncName        string
	InstructionData json.RawMessage
}

// Fiber Initialize the fiber structure
type Fiber struct {
	Name         string
	Instructions []*Instruction
}

// Instruction Initialize a fiber's instruction
type Instruction struct {
	Package         string
	FuncName        string
	InstructionData interface{}
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
				funcName := instruction.FuncName
				instructionData := reflect.ValueOf(instruction.InstructionData).Elem()
				switch instruction.Package {
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
					fmt.Println("FIBER: This package is not integrated yet: " + instruction.Package)
					continue
				}
			}
		}
		fmt.Println("| Fiber Finished |")
		running = 0
	}
}
