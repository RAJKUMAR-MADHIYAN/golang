package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bicycle struct to store details
type Bicycle struct {
	Name          string
	Brand         string
	Price         int
	Features      string
	AfterSales    int // 1 to 5 rating
	Availability  string
	TestRideScore float64
	FinalScore    float64
}

// Function to generate random test ride scores (simulating comfort & handling)
func generateTestRideScore() float64 {
	return 3.5 + rand.Float64()*1.5 // Random score between 3.5 to 5.0
}

// Function to choose the best bicycle based on final scores
func selectBestBicycle(bicycles []Bicycle) Bicycle {
	bestBike := bicycles[0]
	for _, bike := range bicycles {
		if bike.FinalScore > bestBike.FinalScore {
			bestBike = bike
		}
	}
	return bestBike
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Shortlisted bicycles
	bicycles := []Bicycle{
		{"Hero Kids Bike", "Hero", 7000, "Training Wheels, Steel Frame", 4, "Available Online & Stores", generateTestRideScore(), 0},
		{"Btwin Rockrider", "Btwin", 12000, "Mountain Bike, Lightweight Frame", 5, "In Stock Online", generateTestRideScore(), 0},
		{"Firefox Racer", "Firefox", 9500, "BMX, Alloy Frame", 3, "Limited Availability", generateTestRideScore(), 0},
	}

	// Calculate final score (Weighted: AfterSales 40%, TestRide 60%)
	for i := range bicycles {
		bicycles[i].FinalScore = (float64(bicycles[i].AfterSales) * 0.4) + (bicycles[i].TestRideScore * 0.6)
	}

	// Display shortlisted bicycles
	fmt.Println("📌 Shortlisted Bicycles for Final Decision:")
	for _, bike := range bicycles {
		fmt.Printf("🔹 %s | Brand: %s | Price: ₹%d | Features: %s | After-Sales Service: %d⭐ | Availability: %s | Test Ride Score: %.2f⭐ | Final Score: %.2f\n",
			bike.Name, bike.Brand, bike.Price, bike.Features, bike.AfterSales, bike.Availability, bike.TestRideScore, bike.FinalScore)
	}

	// Select the best bicycle
	bestBike := selectBestBicycle(bicycles)

	// Display final recommended bicycle
	fmt.Println("\n✅ Best Bicycle to Buy:")
	fmt.Printf("🏆 %s | Brand: %s | Price: ₹%d | Features: %s | Final Score: %.2f⭐\n",
		bestBike.Name, bestBike.Brand, bestBike.Price, bestBike.Features, bestBike.FinalScore)

	// Accessories & Maintenance
	fmt.Println("\n🛍️ Accessories to Purchase:")
	fmt.Println("✅ Helmet")
	fmt.Println("✅ Knee Pads")
	fmt.Println("✅ Water Bottle Holder")
	fmt.Println("🔧 Look for maintenance package or free servicing.")
}
 