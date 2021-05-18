package mouse

// ClickDatabinder Define the MouseClick parameters
type ClickDatabinder struct {
	MouseButtonName string
}

// ButtonDatabinder Define the MouseButton parameters
type ButtonDatabinder struct {
	ID   string
	Name string
}

// Buttons Define the possibles values of MouseButton
func Buttons() []*ButtonDatabinder {
	return []*ButtonDatabinder{
		{"left", "Left"},
		{"right", "Right"},
	}
}

// ScrollDatabinder Define the MouseScroll parameters
type ScrollDatabinder struct {
	X        int
	XVarName string
	XIsVar   bool
	Y        int
	YVarName string
	YIsVar   bool
}

// MoveDatabinder Define the MouseMove parameters
type MoveDatabinder struct {
	X        int
	XVarName string
	XIsVar   bool
	Y        int
	YVarName string
	YIsVar   bool
}
