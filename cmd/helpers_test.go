package cmd

import (
	"bytes"
)

// This file contains test helpers used by the other *_test.go files

func execCommand(args ...string) ([]byte, []byte, error) {
	stdOut := bytes.NewBufferString("")
	stdErr := bytes.NewBufferString("")
	RootCmd.SetOut(stdOut)
	RootCmd.SetErr(stdErr)
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	return stdOut.Bytes(), stdErr.Bytes(), err
}
