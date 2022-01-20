package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func setup() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flags

	var envVariables = map[string]string{
		"PWD":            "/home/username/Downloads",
		"USER":           "pi",
		"SSH_CONNECTION": "192.168.1.21 54599 192.168.1.31 22",
	}
	for key, value := range envVariables {
		os.Setenv(key, value)
	}
	os.Args = []string{"./bscp", "image.jpeg"}
}

func TestSuccessful(t *testing.T) {
	expectedOutput := "scp -P 22 pi@192.168.1.31:/home/username/Downloads/image.jpeg .\n"
	setup()
	out := CaptureStdout(main)
	if out != expectedOutput {
		t.Errorf("expected %s, but got %s", expectedOutput, out)
	}
}

func TestFailingSshString(t *testing.T) {
	setup()
	os.Setenv("SSH_CONNECTION", "192.168.1.21 54599")
	err := run()
	if !strings.Contains(err.Error(), "SSH_CONNECTION environment variable expected") {
		t.Errorf("SSH_CONNECTION environment variable shouldn't be approved")
	}
}

func TestInvalidArguments(t *testing.T) {
	for _, test := range []struct {
		Name          string
		Args          []string
		ExpectedError string
	}{
		{
			Name:          "Too many arguments",
			Args:          []string{"./bscp", "test.py", "image.jpeg"},
			ExpectedError: "Too many arguments",
		},
		{
			Name:          "No arguments",
			Args:          []string{"./bscp"},
			ExpectedError: "No arguments given",
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setup()
			os.Args = test.Args
			err := run()
			if err.Error() != test.ExpectedError {
				t.Errorf("expected %s, but got %s", test.ExpectedError, err.Error())
			}
		})
	}
}
