package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func run() error {
	pwd := os.Getenv("PWD")
	user := os.Getenv("USER")
	sshEnvVariable := os.Getenv("SSH_CONNECTION")
	words := strings.Fields(sshEnvVariable)
	if len(words) < 4 {
		return errors.New("SSH_CONNECTION environment variable expected form: <client-ip> <client-port> <server-ip> <server-port>")
	}
	serverIP := words[2]
	serverPort := words[3]

	nArgs := len(os.Args[1:])

	errorMsg := ""
	if nArgs > 1 {
		errorMsg = "Too many arguments"
	} else if nArgs == 0 {
		errorMsg = "No arguments given"
	}

	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	file := os.Args[1]
	command := "scp" + " -P " + serverPort + " " + user + "@" + serverIP + ":" + pwd + "/" + file + " ."

	fmt.Println(command)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
