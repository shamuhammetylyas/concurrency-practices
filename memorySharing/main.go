package main

import (
	"fmt"
	"sync"
	"time"
)

var money = 100
var lock sync.Mutex

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("Stingy done")
}

func spendy() {
	for i := 10; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("Spendy done")
}

func main() {

	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)

}
