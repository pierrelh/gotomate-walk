package app

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
