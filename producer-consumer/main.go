package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order: #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd <= 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making pizaa #%d. It will take %d seconds...\n", pizzaNumber, delay)

		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i int

	//run forever or until we receive a quit notification
	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChannel := <-pizzaMaker.quit:
				//close the channels (closing channels in go is GOLDEN RULE)
				close(pizzaMaker.data)
				close(quitChannel)
				return
			}
		}
	}

}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	//print out a message
	fmt.Println("The pizzeria is open for bussiness")
	fmt.Println("----------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background (on its own goroutine)
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				fmt.Println(i.message)
				fmt.Printf("Order #%d is out for delivery\n", i.pizzaNumber)
			} else {
				fmt.Println(i.message)
				fmt.Println("The customer is really mad")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("*** Error closing channel! ", err)
			}
		}
	}

	//print out the ending message
	fmt.Println("-----------------")
	fmt.Println("Done for the day")

	fmt.Printf("We made %d pizzas, but failed to make %d, with %d attempts in total\n", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		fmt.Println("It was an awful day...")
	case pizzasFailed >= 6:
		fmt.Println("It was not a very good day...")
	case pizzasFailed >= 4:
		fmt.Println("It was an OKAY day...")
	case pizzasFailed >= 2:
		fmt.Println("It was a pretty good day...")
	default:
		fmt.Println("It was a great day")
	}
}
