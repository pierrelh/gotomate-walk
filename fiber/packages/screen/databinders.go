package screen

// MouseColorDatabinder Define the GetMouseColor databinder
type MouseColorDatabinder struct {
	Output string
}

// PixelColorDatabinder Define the GetPixelColor databinder
type PixelColorDatabinder struct {
	X        int
	XVarName string
	XIsVar   bool
	Y        int
	YVarName string
	YIsVar   bool
	Output   string
}

// SizeDatabinder Define the GetScreenSize databinder
type SizeDatabinder struct {
	HeightOutput string
	WidthOutput  string
}

// SaveCaptureDatabinder Define theSaveCapture databinder
type SaveCaptureDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
}
