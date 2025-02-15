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
	pathEnv:=os.Getenv("PATH")
	for {
		fmt.Fprint(os.Stdout, "$ ")
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("error in reading from the os.Stdin ", err)
			return
		}
		str = strings.Trim(str, " \r\n\t")

		commands := strings.Split(str, " ")
		/* Enter the execCommand function code here if error comming */

		resp := execCommand(commands,pathEnv)
		fmt.Fprint(os.Stdout, resp)
	}
}
