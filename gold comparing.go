package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GoldStore struct to store gold rate and making charges
type GoldStore struct {
	Name         string
	GoldRate22K  float64
	GoldRate24K  float64
	MakingCharge float64 // Per gram
	DiscountOnMC float64 // Discount percentage on making charge
}

// Jewelry struct to store purchase details
type Jewelry struct {
	Name   string
	Weight float64
	Purity string
}

// Function to generate random gold rates for stores
func generateGoldRates() []GoldStore {
	rand.Seed(time.Now().UnixNano())

	stores := []GoldStore{
		{"Tanishq", 5800 + rand.Float64()*200, 6300 + rand.Float64()*200, 500, 10},
		{"Malabar", 5750 + rand.Float64()*200, 6250 + rand.Float64()*200, 450, 15},
		{"Kalyan", 5700 + rand.Float64()*200, 6200 + rand.Float64()*200, 400, 12},
		{"PC Jeweller", 5850 + rand.Float64()*200, 6350 + rand.Float64()*200, 550, 8},
		{"Joy Alukkas", 5780 + rand.Float64()*200, 6280 + rand.Float64()*200, 480, 10},
	}
	return stores
}

// Function to calculate total price for a jewelry item at a given store
func calculateTotalCost(jewelry Jewelry, store GoldStore) float64 {
	var goldRate float64

	// Select appropriate gold rate based on purity
	if jewelry.Purity == "22K" {
		goldRate = store.GoldRate22K
	} else {
		goldRate = store.GoldRate24K
	}

	// Apply discount on making charges
	discountedMakingCharge := store.MakingCharge * (1 - store.DiscountOnMC/100)

	// Total cost calculation
	totalCost := (goldRate * jewelry.Weight) + (discountedMakingCharge * jewelry.Weight)
	totalCost += totalCost * 0.03 // Adding 3% GST

	return totalCost
}

// Function to find the best store for each jewelry item
func findBestStore(jewelry Jewelry, stores []GoldStore) (GoldStore, float64) {
	bestStore := stores[0]
	lowestPrice := calculateTotalCost(jewelry, stores[0])

	for _, store := range stores[1:] {
		price := calculateTotalCost(jewelry, store)
		if price < lowestPrice {
			lowestPrice = price
			bestStore = store
		}
	}
	return bestStore, lowestPrice
}

func main() {
	// Generate random gold rates
	stores := generateGoldRates()

	// Define jewelry items to buy
	jewelryItems := []Jewelry{
		{"Gold Ring", 5, "22K"},
		{"Gold Bracelet", 8, "22K"},
		{"Gold Bangle", 15, "22K"},
	}

	fmt.Println("📌 Comparing Gold Prices Across Stores...\n")

	// Find the best store for each jewelry item
	for _, item := range jewelryItems {
		bestStore, bestPrice := findBestStore(item, stores)
		fmt.Printf("✅ Best store for %s (%.1f grams, %s): %s\n", item.Name, item.Weight, item.Purity, bestStore.Name)
		fmt.Printf("   - Gold Rate: ₹%.2f/gm\n   - Making Charge: ₹%.2f/gm (after %.0f%% discount)\n   - Total Price: ₹%.2f (incl. GST)\n\n",
			bestStore.GoldRate22K, bestStore.MakingCharge*(1-bestStore.DiscountOnMC/100), bestStore.DiscountOnMC, bestPrice)
	}

	fmt.Println("🎯 Ensure BIS Hallmark Certification & Collect GST Bill before purchase! ✅")
}
