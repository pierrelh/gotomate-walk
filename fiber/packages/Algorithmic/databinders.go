package algorithmic

// Inputs Define the possibles values of KeyboardInput
// func Comparators() []*ComparatorDatabinder {
// 	return []*ComparatorDatabinder{
// 		{""},
// 		{"=="},
// 		{"!="},
// 		{">"},
// 		{">="},
// 		{"<"},
// 		{"<="},
// 	}
// }

// ComparatorDatabinder Define the KeyboardInput parameters
// type ComparatorDatabinder struct {
// 	Name string
// }

// DefineIntDatabinder Define the DefineInt parameters
type DefineIntDatabinder struct {
	Name  string
	Value int
}

// DefineStringDatabinder Define the DefineString parameters
type DefineStringDatabinder struct {
	Name  string
	Value string
}

// DefineBoolDatabinder Define the DefineBool parameters
type DefineBoolDatabinder struct {
	Name  string
	Value bool
}

// IfDatabinder Define the If parameters
// type IfDatabinder struct {
// 	ValueOne         string
// 	Comparator       string
// 	ValueTwo         string
// 	FalseInstruction int
// }
