package app

//Sleep Define the Sleep parameters
type Sleep struct {
	Duration float64
}

//MilliSleep Define the MilliSleep parameters
type MilliSleep struct {
	Duration float64
}

type MouseClick struct {
	MouseButtonName string
}

type MouseButton struct {
	ID   string
	Name string
}

func MouseButtons() []*MouseButton {
	return []*MouseButton{
		{"left", "Left"},
		{"right", "Right"},
	}
}
