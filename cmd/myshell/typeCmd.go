package main

import (
	"fmt"
	"os"
	"strings"
)

func typeCommand(commands []string, pathEnv string) {
	cmdStr := strings.Join(commands[1:], "")

	_, exits := allCommands[cmdStr]
	if exits {
		resp := fmt.Sprintf("%s %s\n", cmdStr, "is a shell builtin")
		fmt.Fprint(os.Stdout, resp)
	} else {
		paths := strings.Split(pathEnv, ":")
		for _, path := range paths {
			str := path + "/" + cmdStr
			_, err := os.Stat(str)
			if err == nil {
				resp := fmt.Sprintf("%s is %s\n", cmdStr, str)
				fmt.Fprint(os.Stdout, resp)
				return
			}
		}
		resp := fmt.Sprintf("%s: %s\n", cmdStr, "not found")
		fmt.Fprint(os.Stdout, resp)
	}

}
