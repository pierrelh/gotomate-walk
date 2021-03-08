package input

import "fmt"

// String Wait for user to set a string
func String(finished chan bool, msg string) string {
	fmt.Println("FIBER: Waiting for user input ...")
	fmt.Println(msg)
	var input string
	fmt.Scanln(&input)
	finished <- true
	return input
}

// Int Wait for user to set a int
func Int(finished chan bool, msg string) int {
	fmt.Println("FIBER: Waiting for user input ...")
	fmt.Println(msg)
	var input int
	fmt.Scanln(&input)
	finished <- true
	return input
}

// Bool Wait for user to set a bool
func Bool(finished chan bool, msg string) bool {
	fmt.Println("FIBER: Waiting for user input ...")
	fmt.Println(msg)
	var input bool
	fmt.Scanln(&input)
	finished <- true
	return input
}
