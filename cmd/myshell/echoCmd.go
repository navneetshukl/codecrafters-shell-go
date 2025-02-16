package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCmd(commands []string) {
	resp := strings.Join(commands[1:], " ")
	resp = fmt.Sprintf("%s\n", resp)

	ch1 := resp[0]
	if string(ch1) == "'" || string(ch1) == `"` {
		resp = resp[1:]
	}
	if string(resp[len(resp)-2]) == "'" || string(resp[len(resp)-2]) == `""` {
		resp1 := ""
		for idx, val := range resp {
			if idx == len(resp)-2 && (string(val) == "'" || string(val) == `"`) {
				continue
			} else {
				resp1 += string(val)
			}
		}
		resp = resp1
	}
	fmt.Fprint(os.Stdout, resp)
}
