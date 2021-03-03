package input

// StringDatabinder Define the String parameters
type StringDatabinder struct {
	Message string
	Var     string
	Value   string
}

// IntDatabinder Define the Int parameters
type IntDatabinder struct {
	Message string
	Var     string
	Value   int
}

// BoolDatabinder Define the Bool parameters
type BoolDatabinder struct {
	Message string
	Var     string
	Value   bool
}
