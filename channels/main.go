package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	name := <-c
	time.Sleep(time.Second * 2)
	fmt.Println("Hello", name)

}

func main() {
	c := make(chan string)

	go greet(c)

	c <- "world"

	fmt.Println("cykysh")
	close(c)
}
