package main

import (
	"fmt"
	"sync"
)

//WaitGroup we mutex-ler ulanylanda bir funksiya iberilen copy edilenok
//yagny bir funksiya ugradylanda shonyn pointeri ugradylyar
//yada package level declare etsen bolyar
//mutex-lerin kopyasy ugradylsa hersi ayratyn mutex hasaplanyar,
//shonun ucin race condition meselani cozenok
//waitGroup hem sholar yaly

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)

}
