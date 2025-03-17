package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RoundResult struct {
	PeopleCrossed int
	TimeTaken     int
}

func simulateRound(round int, results chan<- RoundResult, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as done when function exits

	peopleCrossed := rand.Intn(20) + 1 // Random people crossed (1 to 20)
	timeTaken := rand.Intn(10) + 5     // Random time taken (5 to 14 seconds)

	results <- RoundResult{PeopleCrossed: peopleCrossed, TimeTaken: timeTaken}
	fmt.Printf("Round %d: Crossed %d people, Time taken %d seconds\n", round, peopleCrossed, timeTaken)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	results := make(chan RoundResult, 5) // Buffered channel to collect results
	var wg sync.WaitGroup                // WaitGroup to synchronize goroutines

	// Start 5 goroutines for 5 rounds
	for round := 1; round <= 5; round++ {
		wg.Add(1)
		go simulateRound(round, results, &wg)
	}

	wg.Wait()      // Wait for all goroutines to finish
	close(results) // Close channel after all goroutines complete

	totalPeopleCrossed := 0
	totalTimeTaken := 0

	// Collect results from channel
	for result := range results {
		totalPeopleCrossed += result.PeopleCrossed
		totalTimeTaken += result.TimeTaken
	}

	fmt.Printf("\nTotal: Crossed %d people, Total time taken %d seconds\n", totalPeopleCrossed, totalTimeTaken)
}
