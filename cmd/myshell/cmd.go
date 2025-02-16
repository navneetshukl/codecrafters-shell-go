package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var allCommands = map[string]int{"echo": 0, "exit": 1, "type": 2}

func execCommand(commands []string, pathEnv string) {
	var resp string
	//log.Println("The path env is ", pathEnv)
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		resp = strings.Join(commands[1:], " ")
		resp = fmt.Sprintf("%s\n", resp)
		fmt.Fprint(os.Stdout, resp)

	case "type":
		cmdStr := strings.Join(commands[1:], "")

		_, exits := allCommands[cmdStr]
		if exits {
			resp = fmt.Sprintf("%s %s\n", cmdStr, "is a shell builtin")
			fmt.Fprint(os.Stdout, resp)
		} else {
			paths := strings.Split(pathEnv, ":")
			for _, path := range paths {
				str := path + "/" + cmdStr
				_, err := os.Stat(str)
				if err == nil {
					resp = fmt.Sprintf("%s is %s\n", cmdStr, str)
					fmt.Fprint(os.Stdout, resp)
					return
				}
			}
			resp = fmt.Sprintf("%s: %s\n", cmdStr, "not found")
			fmt.Fprint(os.Stdout, resp)
		}

	case "pwd":
		dir,err:=os.Getwd()
		if err!=nil{
			log.Println("error in getting the current working directory ",err)
			return
		}
		fmt.Fprintf(os.Stdout,"%s\n",dir)
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
