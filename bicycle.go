package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bicycle struct to store details
type Bicycle struct {
	Name           string
	Size           int    // Wheel size in inches
	Type           string // BMX, Mountain, Balance, Training wheels
	Price          int    // Price in INR
	SafetyRating   float64
	CustomerRating float64
	OverallRating  float64
}

// Function to generate random ratings (simulating customer opinions)
func generateRating() float64 {
	return 3.0 + rand.Float64()*2.0 // Random rating between 3.0 to 5.0
}

// Function to compare bicycles and find the best one
func findBestBicycle(bicycles []Bicycle) Bicycle {
	bestBicycle := bicycles[0]

	for _, bike := range bicycles {
		// Higher overall rating is preferred
		if bike.OverallRating > bestBicycle.OverallRating {
			bestBicycle = bike
		}
	}

	return bestBicycle
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	// Sample bicycle data
	bicycles := []Bicycle{
		{"Hero Kids Cycle", 16, "Training Wheels", 7000, generateRating(), generateRating(), 0},
		{"Btwin 500", 20, "Mountain", 12000, generateRating(), generateRating(), 0},
		{"Firefox BMX", 18, "BMX", 9500, generateRating(), generateRating(), 0},
		{"Atlas Balance Bike", 14, "Balance", 5000, generateRating(), generateRating(), 0},
	}

	// Calculate overall rating (Average of Safety & Customer Rating)
	for i := range bicycles {
		bicycles[i].OverallRating = (bicycles[i].SafetyRating + bicycles[i].CustomerRating) / 2
	}

	// Display bicycle options
	fmt.Println("Available Bicycles:")
	for _, bike := range bicycles {
		fmt.Printf("🔹 %s | Size: %d-inch | Type: %s | Price: ₹%d | Safety: %.1f⭐ | Customer: %.1f⭐ | Overall: %.1f⭐\n",
			bike.Name, bike.Size, bike.Type, bike.Price, bike.SafetyRating, bike.CustomerRating, bike.OverallRating)
	}

	// Find the best bicycle based on ratings
	bestBike := findBestBicycle(bicycles)

	// Display the recommended bicycle
	fmt.Println("\n✅ Recommended Bicycle:")
	fmt.Printf("🏆 %s | Size: %d-inch | Type: %s | Price: ₹%d | Overall Rating: %.1f⭐\n",
		bestBike.Name, bestBike.Size, bestBike.Type, bestBike.Price, bestBike.OverallRating)
}
