package packages

import (
	"gotomate/packages/battery"
	"time"

	"gopkg.in/toast.v1"
)

//Sleep Define the Sleep parameters
type Sleep struct {
	Duration float64
}

//MilliSleep Define the MilliSleep parameters
type MilliSleep struct {
	Duration float64
}

//MouseClick Define the MouseClick parameters
type MouseClick struct {
	MouseButtonName string
}

//MouseButton Define the MouseButton parameters
type MouseButton struct {
	ID   string
	Name string
}

//MouseButtons Define the possibles values of MouseButton
func MouseButtons() []*MouseButton {
	return []*MouseButton{
		{"left", "Left"},
		{"right", "Right"},
	}
}

//MouseScroll Define the MouseScroll parameters
type MouseScroll struct {
	X int
	Y int
}

//MouseMove Define the MouseMove parameters
type MouseMove struct {
	X int
	Y int
}

//KeyboardInput Define the KeyboardInput parameters
type KeyboardInput struct {
	Name string
}

//KeyboardInputs Define the possibles values of KeyboardInput
func KeyboardInputs() []*KeyboardInput {
	return []*KeyboardInput{
		{""},
		{"alt"},
		{"cmd"},
		{"shift"},
		{"ctrl"},
		{"enter"},
	}
}

//KeyboardTap Define the KeyboardTap parameters
type KeyboardTap struct {
	Input    string
	Special1 string
	Special2 string
}

//ClipboardWrite Define the ClipboardWrite parameters
type ClipboardWrite struct {
	Content string
}

//ClipboardRead Define the ClipboardRead parameters
type ClipboardRead struct {
	Var   string
	Value string
}

//LogPrint Define the LogPrint parameters
type LogPrint struct {
	Log string
}

//NotificationCreate Define the NotificationCreate parameters
type NotificationCreate struct {
	Title   string
	Message string
	Actions []toast.Action
}

//UserBattery Define the UserBattery parameters
type UserBattery struct {
	Var   string
	Value *battery.Battery
}

//BatteryState Define the BatteryState parameters
type BatteryState struct {
	BatteryName string
	Var         string
	Value       battery.State
}

//BatteryPercentage Define the BatteryPercentage parameters
type BatteryPercentage struct {
	BatteryName string
	Var         string
	Value       float64
}

//BatteryRemainingTime Define the BatteryRemainingTime parameters
type BatteryRemainingTime struct {
	BatteryName string
	Var         string
	Value       time.Duration
}

//BatteryChargeRate Define the BatteryChargeRate parameters
type BatteryChargeRate struct {
	BatteryName string
	Var         string
	Value       time.Duration
}

//BatteryCurrentCapacity Define the BatteryCurrentCapacity parameters
type BatteryCurrentCapacity struct {
	BatteryName string
	Var         string
	Value       float64
}

//BatteryLastFullCapacity Define the BatteryLastFullCapacity parameters
type BatteryLastFullCapacity struct {
	BatteryName string
	Var         string
	Value       float64
}

//BatteryDesignCapacity Define the BatteryDesignCapacity parameters
type BatteryDesignCapacity struct {
	BatteryName string
	Var         string
	Value       float64
}

//BatteryVoltage Define the BatteryVoltage parameters
type BatteryVoltage struct {
	BatteryName string
	Var         string
	Value       float64
}

//BatteryDesignVoltage Define the BatteryDesignVoltage parameters
type BatteryDesignVoltage struct {
	BatteryName string
	Var         string
	Value       float64
}

//SysClock Define the SysClock parameters
type SysClock struct {
	Var   string
	Value [3]int
}

//SysTime Define the SysTime parameters
type SysTime struct {
	Var   string
	Value time.Time
}
