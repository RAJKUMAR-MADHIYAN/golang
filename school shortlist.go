package main

import (
	"fmt"
	"sort"
)

// School struct to store school details
type School struct {
	Name             string
	Curriculum       string
	AnnualFees       int
	StudentTeacherRatio float64
	Facilities       map[string]bool // Facilities like Sports, Labs, Transport, etc.
	OverallScore     float64
}

// Function to calculate a weighted score for school ranking
func calculateScore(s School) float64 {
	score := 0.0

	// Lower fees get higher score
	if s.AnnualFees <= 200000 {
		score += 10
	} else if s.AnnualFees <= 300000 {
		score += 7
	} else {
		score += 5
	}

	// Lower student-teacher ratio gets a better score
	if s.StudentTeacherRatio <= 20 {
		score += 10
	} else if s.StudentTeacherRatio <= 30 {
		score += 7
	} else {
		score += 5
	}

	// More facilities add to the score
	for _, available := range s.Facilities {
		if available {
			score += 2
		}
	}

	return score
}

// Function to find the best school
func findBestSchool(schools []School) School {
	sort.Slice(schools, func(i, j int) bool {
		return schools[i].OverallScore > schools[j].OverallScore
	})
	return schools[0]
}

func main() {
	// Define shortlisted schools
	schools := []School{
		{"Greenwood International", "IB", 280000, 18, map[string]bool{"Sports": true, "Labs": true, "Transport": true, "Extracurricular": true}, 0},
		{"Oakridge International", "IGCSE", 320000, 22, map[string]bool{"Sports": true, "Labs": true, "Transport": false, "Extracurricular": true}, 0},
		{"The International School", "CBSE International", 250000, 20, map[string]bool{"Sports": true, "Labs": false, "Transport": true, "Extracurricular": true}, 0},
		{"DPS Global", "ICSE", 220000, 25, map[string]bool{"Sports": false, "Labs": true, "Transport": true, "Extracurricular": true}, 0},
		{"Indus International", "IB", 350000, 15, map[string]bool{"Sports": true, "Labs": true, "Transport": true, "Extracurricular": true}, 0},
	}

	// Calculate scores for each school
	for i := range schools {
		schools[i].OverallScore = calculateScore(schools[i])
	}

	// Find the best school
	bestSchool := findBestSchool(schools)

	// Display school rankings
	fmt.Println("📌 School Comparison Results:\n")
	for _, s := range schools {
		fmt.Printf("🏫 %s (%s) - Fees: ₹%d, Ratio: %.1f, Score: %.2f\n",
			s.Name, s.Curriculum, s.AnnualFees, s.StudentTeacherRatio, s.OverallScore)
	}

	// Final Decision
	fmt.Printf("\n🎯 Best school for admission: %s (%s)\n", bestSchool.Name, bestSchool.Curriculum)
	fmt.Println("🔍 Ensure to visit & review before finalizing!")
}
