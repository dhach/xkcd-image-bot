package helpers

import (
	"fmt"
	"os"
)

// PrintErrorAnExit prints an error an then exists the programm
// The exit code will always be 1
func PrintErrorAndExit(errorMessage error) {
	fmt.Println("[E] There was an error:")
	fmt.Println("[E]", errorMessage)
	os.Exit(1)
}
