package packages

import (
	"fmt"
	"gotomate/fiber"
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
)

// PackageDecode return the struct to use to decode an unmarshalled data
func PackageDecode(instruction *fiber.LoadingInstruction) interface{} {
	switch instruction.Package {
	case "Battery":
		return battery.Decode(instruction.FuncName)
	case "Clipboard":
		return clipboard.Decode(instruction.FuncName)
	case "Flow":
		return flow.Decode(instruction.FuncName)
	case "Input":
		return input.Decode(instruction.FuncName)
	case "Keyboard":
		return keyboard.Decode(instruction.FuncName)
	case "Log":
		return log.Decode(instruction.FuncName)
	case "Mouse":
		return mouse.Decode(instruction.FuncName)
	case "Notification":
		return notification.Decode(instruction.FuncName)
	case "Process":
		return process.Decode(instruction.FuncName)
	case "Screen":
		return screen.Decode(instruction.FuncName)
	case "Sleep":
		return sleep.Decode(instruction.FuncName)
	case "Systime":
		return systime.Decode(instruction.FuncName)
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the package to decode the instruction")
		return nil
	}
}
