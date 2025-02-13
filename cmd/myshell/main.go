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

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("error in reading from the os.Stdin ", err)
			return
		}
		str = strings.Trim(str, " \r\n\t")
		if str=="exit 0"{
			os.Exit(0)
		}
		resp := fmt.Sprintf("%s: %s", str, "command not found \n")
		fmt.Fprint(os.Stdout, resp)
	}
}
