package automate

var Packages = []string{
	"Log",
	"Battery",
	"Clipboard",
	"Sleep",
	"Notification",
}

var SubPackages = map[string][]string{
	"Log": {
		"log1",
		"log2",
	},
	"Battery": {
		"battery1",
		"battery2",
	},
	"Clipboard": {
		"clipboard1",
		"clipboard2",
	},
	"Sleep": {
		"sleep1",
		"sleep2",
	},
	"Notification": {
		"notification1",
		"notification2",
	},
}
