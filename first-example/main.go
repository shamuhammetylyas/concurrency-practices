package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
	}

	for i, x := range numbers {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Text 2", &wg)
}
