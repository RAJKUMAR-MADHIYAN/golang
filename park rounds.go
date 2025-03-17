package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	totalPeopleCrossed := 0
	totalTimeTaken := 0

	for round := 1; round <= 5; round++ {
		peopleCrossed := rand.Intn(20) + 1 // Random people crossed (1 to 20)
		timeTaken := rand.Intn(10) + 5     // Random time taken (5 to 14 seconds)

		totalPeopleCrossed += peopleCrossed
		totalTimeTaken += timeTaken

		fmt.Printf("Round %d: Crossed %d people, Time taken %d seconds\n", round, peopleCrossed, timeTaken)
	}

	fmt.Printf("\nTotal: Crossed %d people, Total time taken %d seconds\n", totalPeopleCrossed, totalTimeTaken)
}
