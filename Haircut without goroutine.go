package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BarberShop struct {
	capacity     int
	waitingQueue []int
}

func NewBarberShop(capacity int) *BarberShop {
	return &BarberShop{
		capacity:     capacity,
		waitingQueue: []int{},
	}
}

func (shop *BarberShop) Barber() {
	for len(shop.waitingQueue) > 0 {
		customerID := shop.waitingQueue[0]
		shop.waitingQueue = shop.waitingQueue[1:]
		fmt.Printf("Barber is cutting hair for Customer %d\n", customerID)
		time.Sleep(2 * time.Second) // Simulate haircut time
		fmt.Printf("Customer %d is done with haircut\n", customerID)
	}
}

func (shop *BarberShop) Customer(customerID int) {
	if len(shop.waitingQueue) < shop.capacity {
		shop.waitingQueue = append(shop.waitingQueue, customerID)
		fmt.Printf("Customer %d is waiting\n", customerID)
	} else {
		fmt.Printf("Customer %d found no seat and left\n", customerID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	shop := NewBarberShop(5) // Shop with 3 waiting seats

	numCustomers := 10
	for i := 1; i <= numCustomers; i++ {
		shop.Customer(i)
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // Random arrival
	}

	shop.Barber()
	fmt.Println("Shop is closed!")
}
