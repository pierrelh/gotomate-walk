package app

import "gopkg.in/toast.v1"

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
	Key     string
	Content string
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
