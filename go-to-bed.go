package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, nocturnal being! Tell me, what's keeping you up? Coding? Goding? Reading?")

	for {
		fmt.Print("Enter your reason or 'quit' to exit: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "quit" || text == "I'm going to bed now" {
			break
		}

		// Based on the user's input, provide different responses.
		// Use if/else if/else statements to compare the text to various strings,
		// and print out an appropriate response using fmt.Println().

		if strings.Contains(text, "coding") {
			fmt.Println("Okay, forget bed time - it's code time...")
		} else if text == "goding" {
			fmt.Println("Oooh, goding? Okay, forget code time, and definitely forget bed time! It's gode time...")
		} else if text == "reading" {
			fmt.Println("Ah nice! Reading is really good fo - wait, you're reading code aren't you?")
		} else {
			fmt.Println("Go to bed!")
		}
	}
	fmt.Println("Goodnight, sleep tight!")
}
