package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	defer wg.Done()
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello world"

	wg.Add(1)
	updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}
