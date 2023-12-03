package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Order int

func (o Order) String() string {
	return fmt.Sprintf("order-%02d", o)
}

func processOrder(_ Order) {
	processingTime := time.Duration(2+rand.Intn(2)) * time.Second
	time.Sleep(processingTime)
}

func waitForOrders() {
	processingTime := time.Duration(rand.Intn(2)) * time.Second
	time.Sleep(processingTime)
}

func worker(name string, ch <-chan Order) {
	for o := range ch {
		log.Printf("%s: I got %v, I will process it\n", name, o)
		processOrder(o)
		log.Printf("%s: I completed %v, I'm ready to take a new order\n", name, o)
	}
	log.Printf("%s: I'm done\n", name)
}

func main() {
	ch := make(chan Order, 3)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		worker("Candier", ch)
	}()

	go func() {
		defer wg.Done()
		worker("Stringer", ch)
	}()

	for i := 0; i < 10; i++ {
		waitForOrders()
		o := Order(i)
		log.Printf("Partier: I %v, I will pass it to the channel\n", o)
		ch <- o
	}

	log.Println("No more orders, closing the channel to signify workers to stop")
	close(ch)

	log.Println("Wait for workers to gracefully stop")
	wg.Wait()

	log.Println("All done")
}
