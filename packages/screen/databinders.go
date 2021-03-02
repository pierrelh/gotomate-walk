package screen

// PixelColorDatabinder Define the GetPixelColor databinder
type PixelColorDatabinder struct {
	X     int
	Y     int
	Var   string
	Value string
}

// MouseColorDatabinder Define the GetMouseColor databinder
type MouseColorDatabinder struct {
	Var   string
	Value string
}

// SizeDatabinder Define the GetScreenSize databinder
type SizeDatabinder struct {
	Var    string
	Height string
	Width  string
}

// SaveCaptureDatabinder Define the SaveCapture databinder
type SaveCaptureDatabinder struct {
	Path string
}
