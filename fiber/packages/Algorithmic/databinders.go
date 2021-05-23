package algorithmic

// Comparators Define the possibles values of algorithmic comparators
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

// ForDatabinder Define the For parameters
type ForDatabinder struct {
	VarOneName              string
	Comparator              string
	ValueTwo                int
	VarTwoName              string
	TwoIsVar                bool
	IncrementVarName        string
	Increment               int
	IncrementIsVar          bool
	FalseInstruction        int
	FalseInstructionVarName string
	FalseInstructionIsVar   bool
}

// IfDatabinder Define the If parameters
type IfDatabinder struct {
	ValueOne                interface{}
	VarOneName              string
	OneIsVar                bool
	Comparator              string
	ValueTwo                interface{}
	VarTwoName              string
	TwoIsVar                bool
	FalseInstruction        int
	FalseInstructionVarName string
	FalseInstructionIsVar   bool
}
