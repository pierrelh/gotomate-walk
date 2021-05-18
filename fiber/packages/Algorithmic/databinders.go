package algorithmic

// Inputs Define the possibles values of KeyboardInput
func Comparators() []*ComparatorDatabinder {
	return []*ComparatorDatabinder{
		{""},
		{"=="},
		{"!="},
		{">"},
		{">="},
		{"<"},
		{"<="},
	}
}

// ComparatorDatabinder Define the Comparators parameters
type ComparatorDatabinder struct {
	Name string
}

// IfDatabinder Define the If parameters
type IfDatabinder struct {
	ValueOne         interface{}
	VarOneName       string
	OneIsVar         bool
	Comparator       string
	ValueTwo         interface{}
	VarTwoName       string
	TwoIsVar         bool
	FalseInstruction int
}
