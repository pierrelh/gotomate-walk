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
	X      int
	XIsVar bool
	Y      int
	YIsVar bool
}

// MoveDatabinder Define the MouseMove parameters
type MoveDatabinder struct {
	X      int
	XIsVar bool
	Y      int
	YIsVar bool
}
