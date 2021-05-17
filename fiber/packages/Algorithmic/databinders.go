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

// DefineBoolDatabinder Define the DefineBool parameters
type DefineBoolDatabinder struct {
	Name  string
	Value bool
	IsVar bool
}

// DefineIntDatabinder Define the DefineInt parameters
type DefineIntDatabinder struct {
	Name  string
	Value int
	IsVar bool
}

// DefineStringDatabinder Define the DefineString parameters
type DefineStringDatabinder struct {
	Name  string
	Value string
	IsVar bool
}

// IfDatabinder Define the If parameters
type IfDatabinder struct {
	ValueOne         interface{}
	OneIsVar         bool
	Comparator       string
	ValueTwo         interface{}
	TwoIsVar         bool
	FalseInstruction int
}
