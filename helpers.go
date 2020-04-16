package main

import (
	"fmt"
	"os"
)

// PrintErrorAndExit takes a custom message and error, prints it and then exits the program
func PrintErrorAndExit(customMessage string, errorMessage error) {
	fmt.Println("[E] Error ", customMessage, ":")
	fmt.Println(errorMessage)
	os.Exit(1)
}
