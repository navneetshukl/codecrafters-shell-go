package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func init() {
	log.SetPrefix("trace: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func main() {

	TYPE := []string{"echo", "exit", "type"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("error in reading from the os.Stdin ", err)
			return
		}
		str = strings.Trim(str, " \r\n\t")

		fmt.Println("Strrrr is ",str)

		commands := strings.Split(str, " ")

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

		fmt.Fprint(os.Stdout, resp)
	}
}
