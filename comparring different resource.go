package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bicycle struct to store details
type Bicycle struct {
	Name         string
	Size         int
	Type         string
	Price        int
	Brand        string
	SafetyRating float64
	OnlineRating float64
	StoreRating  float64
	FinalScore   float64
}

// Function to generate a random rating between 3.0 and 5.0
func generateRating() float64 {
	return 3.0 + rand.Float64()*2.0
}

// Function to evaluate bicycles and recommend the best one
func findBestBicycle(bicycles []Bicycle) Bicycle {
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

	// Simulated bicycle data from different sources
	bicycles := []Bicycle{
		{"Hero Kids Bike", 16, "Training Wheels", 7000, "Hero", generateRating(), generateRating(), generateRating(), 0},
		{"Btwin Rockrider", 20, "Mountain", 12000, "Btwin", generateRating(), generateRating(), generateRating(), 0},
		{"Firefox Racer", 18, "BMX", 9500, "Firefox", generateRating(), generateRating(), generateRating(), 0},
		{"Atlas Smart Balance", 14, "Balance", 5000, "Atlas", generateRating(), generateRating(), generateRating(), 0},
	}

	// Calculate Final Score (weighted average of safety, online, and store ratings)
	for i := range bicycles {
		bicycles[i].FinalScore = (bicycles[i].SafetyRating*0.4 + bicycles[i].OnlineRating*0.3 + bicycles[i].StoreRating*0.3)
	}

	// Display bicycle options
	fmt.Println("📌 Comparing Bicycles from Different Sources:")
	for _, bike := range bicycles {
		fmt.Printf("🔹 %s | Brand: %s | Size: %d-inch | Type: %s | Price: ₹%d | Safety: %.1f⭐ | Online: %.1f⭐ | Store: %.1f⭐ | Final Score: %.2f\n",
			bike.Name, bike.Brand, bike.Size, bike.Type, bike.Price, bike.SafetyRating, bike.OnlineRating, bike.StoreRating, bike.FinalScore)
	}

	// Find the best bicycle
	bestBike := findBestBicycle(bicycles)

	// Display the final recommendation
	fmt.Println("\n✅ Recommended Bicycle:")
	fmt.Printf("🏆 %s | Brand: %s | Size: %d-inch | Type: %s | Price: ₹%d | Final Score: %.2f⭐\n",
		bestBike.Name, bestBike.Brand, bestBike.Size, bestBike.Type, bestBike.Price, bestBike.FinalScore)
}
