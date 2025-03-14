package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BarberShop struct {
	waitingRoom chan int
	barberDone  chan bool
	wg          sync.WaitGroup
}

func NewBarberShop(capacity int) *BarberShop {
	return &BarberShop{
		waitingRoom: make(chan int, capacity), // Queue for customers
		barberDone:  make(chan bool),
	}
}

func (shop *BarberShop) Barber() {
	for {
		select {
		case customerID := <-shop.waitingRoom:
			fmt.Printf("Barber is cutting hair for Customer %d\n", customerID)
			time.Sleep(2 * time.Second) // Simulate haircut time
			fmt.Printf("Customer %d is done with haircut\n", customerID)
			shop.barberDone <- true
		}
	}
}

func (shop *BarberShop) Customer(customerID int) {
	select {
	case shop.waitingRoom <- customerID:
		fmt.Printf("Customer %d is waiting\n", customerID)
		<-shop.barberDone // Wait for haircut to finish
	default:
		fmt.Printf("Customer %d found no seat and left\n", customerID)
	}
	shop.wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	shop := NewBarberShop(3) // Shop with 3 waiting seats
	go shop.Barber()         // Start barber goroutine

	numCustomers := 10
	shop.wg.Add(numCustomers)
	for i := 1; i <= numCustomers; i++ {
		go shop.Customer(i)
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // Random arrival
	}

	shop.wg.Wait()
	fmt.Println("Shop is closed!")
}
