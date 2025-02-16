package main

import (
	"fmt"
	"os"
)

func cdCmd(commands []string) {
	dirName := commands[1]
	if dirName == "~" {

		_, _ = os.UserHomeDir()

	} else {
		err := os.Chdir(dirName)
		if err != nil {
			fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", dirName)
			return
		}
	}
}
