package algorithmic

import "github.com/lxn/walk/declarative"

// DefineTemplate Dialog's Define Template
var DefineTemplate = []declarative.Widget{
	declarative.Label{
		Text: "Name:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Name"),
		CompactHeight: true,
	},
	declarative.Label{
		Text: "Value:",
	},
	declarative.TextEdit{
		Text:          declarative.Bind("Value"),
		CompactHeight: true,
	},
}

// IfTemplate Dialog's If Template
// var IfTemplate = []declarative.Widget{
// 	declarative.Label{
// 		Text: "Value 1:",
// 	},
// 	declarative.TextEdit{
// 		Text:          declarative.Bind("ValueOne"),
// 		CompactHeight: true,
// 	},
// 	declarative.Label{
// 		Text: "Comparator:",
// 	},
// 	declarative.ComboBox{
// 		Value:         declarative.Bind("Comparator", declarative.SelRequired{}),
// 		BindingMember: "Name",
// 		DisplayMember: "Name",
// 		Model:         Comparators(),
// 	},
// 	declarative.Label{
// 		Text: "Value 2:",
// 	},
// 	declarative.TextEdit{
// 		Text:          declarative.Bind("ValueTwo"),
// 		CompactHeight: true,
// 	},
// 	declarative.Label{
// 		Text: "If false instruction ID:",
// 	},
// 	declarative.TextEdit{
// 		Text:          declarative.Bind("FalseInstruction"),
// 		CompactHeight: true,
// 	},
// }
