package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	print_welcome()

	for {
		print_prompt()
		text := read_buffer(reader)
		if text == ".exit" {
			return
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", text)
		}
	}
}

func print_prompt() {
	fmt.Print("db > ")
}

func print_welcome() {
	fmt.Println("Welcome to gosql")
}

func read_buffer(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return text
}
