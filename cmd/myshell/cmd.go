package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var allCommands = map[string]int{"echo": 0, "exit": 1, "type": 2, "pwd": 3}

func execCommand(commands []string, pathEnv, homeEnv string) {
	var resp string
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		resp = strings.Join(commands[1:], " ")
		resp = fmt.Sprintf("%s\n", resp)

		ch1 := resp[0]
		if string(ch1) == "'" {
			resp = resp[1:]
		}
		if string(resp[len(resp)-2]) == "'" {
			resp1 := ""
			for idx, val := range resp {
				//fmt.Println("index and character is ", idx, string(val))
				if idx == len(resp)-2 && string(val) == "'" {
					continue
				} else {
					resp1 += string(val)
				}
			}
			resp = resp1
		}
		//fmt.Println("Length is ", len(resp))
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
