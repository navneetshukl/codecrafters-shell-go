package main

import (
	"fmt"
	"os"
	"strings"
)

// func echoCmd(commands []string) {
// 	resp := strings.Join(commands[1:], " ")
// 	resp = fmt.Sprintf("%s\n", resp)

// 	ch1 := resp[0]
// 	if string(ch1) == "'" {
// 		resp = resp[1:]
// 	}

// 	if string(resp[len(resp)-2]) == "'" {
// 		resp1 := ""
// 		for idx, val := range resp {
// 			if idx == len(resp)-2 && string(val) == "'" {
// 				continue
// 			} else {
// 				resp1 += string(val)
// 			}
// 		}
// 		resp = resp1
// 	}
// 	fmt.Fprint(os.Stdout, resp)
// }

func echoCmd(commands []string) {
	var parts []string
	inSingleQuotes := false
	var currentPart strings.Builder

	for _, word := range commands[1:] {
		// Handle single quotes
		if strings.HasPrefix(word, "'") {
			inSingleQuotes = true
			word = word[1:] // Remove starting quote
		}
		if strings.HasSuffix(word, "'") {
			inSingleQuotes = false
			word = word[:len(word)-1] // Remove ending quote
		}

		// Add to the current part
		if inSingleQuotes {
			currentPart.WriteString(word + " ")
		} else {
			currentPart.WriteString(word)
			parts = append(parts, currentPart.String())
			currentPart.Reset()
		}
	}

	// Join and print the final result
	resp := strings.Join(parts, " ")
	fmt.Fprintln(os.Stdout, resp)
}
