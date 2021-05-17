package screen

// MouseColorDatabinder Define the GetMouseColor databinder
type MouseColorDatabinder struct {
	Output string
}

// PixelColorDatabinder Define the GetPixelColor databinder
type PixelColorDatabinder struct {
	X      int
	XIsVar bool
	Y      int
	YIsVar bool
	Output string
}

// SizeDatabinder Define the GetScreenSize databinder
type SizeDatabinder struct {
	HeightOutput string
	WidthOutput  string
}

// SaveCaptureDatabinder Define theSaveCapture databinder
type SaveCaptureDatabinder struct {
	Path      string
	PathIsVar bool
}
