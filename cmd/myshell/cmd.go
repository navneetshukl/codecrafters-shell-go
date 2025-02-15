package main

import (
	"fmt"
	"os"
	"strings"
)

func execCommand(commands []string) string {
	TYPE := []string{"echo", "exit", "type"}
	var resp string
	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		resp = strings.Join(commands[1:], " ")
		resp = fmt.Sprintf("%s\n", resp)

	case "type":
		cmdStr := strings.Join(commands[1:], " ")
		isFound := false
		for _, v := range TYPE {
			if v == cmdStr {
				isFound = true
				break
			}
		}
		if isFound {
			resp = fmt.Sprintf("%s %s \n", cmdStr, "is a shell builtin")
		} else {
			resp = fmt.Sprintf("%s: %s \n", cmdStr, "not found")
		}

	default:
		resp = strings.Join(commands, " ")
		resp = fmt.Sprintf("%s: command not found \n", resp)

	}
	return resp
}
