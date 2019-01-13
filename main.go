package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() string {
	// Run our API caller in a go routine when user types 'start'
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// Remove newline
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	fetcher := NewFetcher()
	for {
		text := getInput()
		// Case statement?
		if text == "start" {
			go fetcher.start()
		} else if text == "stop" {
			fetcher.stop()
		} else {
			fmt.Println(fmt.Sprintf("Invalid command: %s", text))
		}
	}
}
