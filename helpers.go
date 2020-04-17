package main

import (
	"fmt"
	"os"
)

// printErrorAndExit takes an error message, prints it and then exits the program
func printErrorAndExit(errorMessage error) {
	fmt.Println("[E] ", errorMessage)
	os.Exit(1)
}
