package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var allCommands = map[string]int{"echo": 0, "exit": 1, "type": 2, "pwd": 3}

func execCommand(commands []string, pathEnv, homeEnv string) {
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":

		echoCmd(commands)

	case "type":

		typeCommand(commands, pathEnv)

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			log.Println("error in getting the current working directory ", err)
			return
		}
		fmt.Fprintf(os.Stdout, "%s\n", dir)
	case "cd":

		cdCmd(commands, homeEnv)

	default:
		command := exec.Command(commands[0], commands[1:]...)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			fmt.Printf("%s: command not found\n", commands[0])
		}
	}
}
