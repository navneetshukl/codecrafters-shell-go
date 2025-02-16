package main

import (
	"fmt"
	"os"
)

func cdCmd(commands []string, homeEnv string) {
	dirName := commands[1]
	if dirName == "~" {
		dirName = homeEnv

	}
	err := os.Chdir(dirName)
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", dirName)
		return
	}
}
