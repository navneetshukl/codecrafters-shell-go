package main

import (
	"fmt"
	"os"
	"strings"
)

var allCommands = map[string]int{"echo": 0, "exit": 1, "type": 2}

func execCommand(commands []string, pathEnv string) string {
	var resp string
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		resp = strings.Join(commands[1:], " ")
		resp = fmt.Sprintf("%s\n", resp)

	case "type":
		cmdStr := strings.Join(commands[1:], "")

		_, exits := allCommands[cmdStr]
		if exits {
			resp = fmt.Sprintf("%s %s \n", cmdStr, "is a shell builtin")
			return resp
		} else {
			paths := strings.Split(pathEnv, ":")
			for _, path := range paths {
				str := path + "/" + cmdStr
				_, err := os.Stat(str)
				if err != nil {
					resp = fmt.Sprintf("%s is %s ", cmdStr, str)
					return resp
				}
			}
		}
		resp = fmt.Sprintf("%s: %s \n", cmdStr, "not found")

	default:
		resp = strings.Join(commands, " ")
		resp = fmt.Sprintf("%s: command not found \n", resp)

	}
	return resp
}
