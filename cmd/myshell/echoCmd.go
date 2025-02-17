package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCmd(commands []string) {
	//resp := strings.Join(commands[1:], " ")
	commands = commands[1:]
	resp := ""
	for _, val := range commands {
		if len(val) > 0 {
			resp += val + " "
		}
	}
	//fmt.Println("Resp is ",resp)
	//fmt.Println("Length is ",len(resp))
	// for idx,val:=range resp{
	// 	fmt.Println("Index and string is ",idx,string(val))
	// }
	isquote:=false

	ch1 := resp[0]
	if string(ch1) == "'" {
		resp = resp[1:]
		isquote=true
	}
	//fmt.Println("char is ",string(resp[len(resp)-2]))
	resp=strings.TrimSpace(resp)

	if isquote {
		resp1 := ""
		for _, val := range resp {
			if string(val) == "'" {
				continue
			} else {
				resp1 += string(val)
			}
		}
		resp = resp1
	}
	resp+="\n"
	fmt.Fprint(os.Stdout, resp)
}

// func echoCmd(commands []string) {
// 	var parts []string
// 	inSingleQuotes := false
// 	var currentPart strings.Builder

// 	for _, word := range commands[1:] {
// 		if strings.HasPrefix(word, "'") {
// 			inSingleQuotes = true
// 			word = word[1:]
// 		}
// 		if strings.HasSuffix(word, "'") {
// 			inSingleQuotes = false
// 			word = word[:len(word)-1]
// 		}

// 		if inSingleQuotes {
// 			currentPart.WriteString(word + " ")
// 		} else {
// 			currentPart.WriteString(word)
// 			parts = append(parts, currentPart.String())
// 			currentPart.Reset()
// 		}
// 	}

// 	resp := strings.Join(parts, " ")
// 	fmt.Fprintln(os.Stdout, resp)
// }
