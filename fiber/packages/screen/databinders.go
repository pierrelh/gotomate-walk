package screen

// PixelColorDatabinder Define the GetPixelColor databinder
type PixelColorDatabinder struct {
	X      int
	Y      int
	Output string
}

// MouseColorDatabinder Define the GetMouseColor databinder
type MouseColorDatabinder struct {
	Output string
}

// SizeDatabinder Define the GetScreenSize databinder
type SizeDatabinder struct {
	Output string
}

// SaveCaptureDatabinder Define the SaveCapture databinder
type SaveCaptureDatabinder struct {
	Path string
}
