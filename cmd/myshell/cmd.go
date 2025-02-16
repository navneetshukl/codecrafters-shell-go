package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var allCommands = map[string]int{"echo": 0, "exit": 1, "type": 2, "pwd": 3}

func execCommand(commands []string, pathEnv string) {
	var resp string
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		resp = strings.Join(commands[1:], " ")
		resp = fmt.Sprintf("%s\n", resp)
		fmt.Fprint(os.Stdout, resp)

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
		// dirName := commands[1]
		// err := os.Chdir(dirName)
		// if err != nil {
		// 	fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", dirName)
		// 	return
		// }
		cdCmd(commands)
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
