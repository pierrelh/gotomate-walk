package conversion

// IntToStringDatabinder Define the IntToString parameters
type IntToStringDatabinder struct {
	Input        int
	InputVarName string
	InputIsVar   bool
	Output       string
}

// StringConversionDatabinder Define the StringConversion parameters
type StringConversionDatabinder struct {
	Input        string
	InputVarName string
	InputIsVar   bool
	Output       string
}
