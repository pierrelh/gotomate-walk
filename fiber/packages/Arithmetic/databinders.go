package arithmetic

// ArithmeticDatabinder Define the Arithmetic parameters
type ArithmeticDatabinder struct {
	ValueOne      string
	ValueOneIsVar bool
	ValueTwo      string
	ValueTwoIsVar bool
	Output        string
}

// SqrtDatabinder Define the Sqrt parameters
type SqrtDatabinder struct {
	Value      string
	ValueIsVar bool
	Output     string
}
