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
	 fmt.Fprint(os.Stdout, "$ ")

	str,err:=bufio.NewReader(os.Stdin).ReadString('\n')
	if err!=nil{
		log.Println("error in reading from the os.Stdin ",err )
		return
	}
	str=strings.Trim(str," \r\n\t")
	resp:=fmt.Sprintf("%s: %s",str,"command not found")
	fmt.Fprint(os.Stdout,resp)
}
