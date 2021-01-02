package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printErrorAndExit(t *testing.T) {
	genericError := errors.New("Generic error")
	if os.Getenv("TEST_CRASH") == "1" {
		printErrorAndExit(genericError)
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_printErrorAndExit")
	cmd.Env = append(os.Environ(), "TEST_CRASH=1")
	err := cmd.Run()
	assert.EqualError(t, err, "exit status 1")
}
