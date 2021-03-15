package flow

import "fmt"

// Read read string from clipboard
func Start(finished chan bool) {
	fmt.Println("| Fiber Start |")
	finished <- true
}

// Write write string to clipboard
func End(finished chan bool) {
	fmt.Println("| Fiber Finished |")
	finished <- true
}
