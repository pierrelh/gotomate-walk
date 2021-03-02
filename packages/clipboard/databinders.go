package clipboard

// WriteDatabinder Define the ClipboardWrite parameters
type WriteDatabinder struct {
	Content string
}

// ReadDatabinder Define the ClipboardRead parameters
type ReadDatabinder struct {
	Var   string
	Value string
}
