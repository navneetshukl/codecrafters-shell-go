package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCmd(commands []string) {
	//resp := strings.Join(commands[1:], " ")
	resp := ""
	//fmt.Println("Commands is ", commands)
	for _, val := range commands {
		//fmt.Println("Length and character is ", len(val), val)
		if len(val) > 0 && val != "echo" {
			resp += val + " "
		}
	}
	for idx := range commands {
		commands[idx] = strings.TrimSpace(commands[idx])
	}
	resp = fmt.Sprintf("%s\n", resp)

	ch1 := resp[0]
	if string(ch1) == "'" {
		resp = resp[1:]
	}

	if string(resp[len(resp)-2]) == "'" {
		resp1 := ""
		for idx, val := range resp {
			if idx == len(resp)-2 && string(val) == "'" {
				continue
			} else {
				resp1 += string(val)
			}
		}
		resp = resp1
	}
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
