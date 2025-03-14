package main

import (
	"fmt"
	"time"
)

type House struct {
	rooms map[string]bool
}

func NewHouse() *House {
	return &House{
		rooms: map[string]bool{
			"Room 1": false,
			"Room 2": false,
			"Hall":   false,
		},
	}
}

func (h *House) CleanRoom(room string) {
	fmt.Printf("Cleaning %s...\n", room)
	time.Sleep(2 * time.Second) // Simulate cleaning time
	h.rooms[room] = true
	fmt.Printf("%s is now shining and clean! ✨\n", room)
}

func (h *House) IsHouseClean() bool {
	for _, clean := range h.rooms {
		if !clean {
			return false
		}
	}
	return true
}

func main() {
	house := NewHouse()
	helper := []string{"Helper 1", "Helper 2"}

	for room := range house.rooms {
		fmt.Printf("%s is starting to clean %s.\n", helper[0], room)
		house.CleanRoom(room)
		helper = append(helper[1:], helper[0]) // Rotate helper
	}

	if house.IsHouseClean() {
		fmt.Println("The entire house is now shining and spotless! ✨🏠✨")
	} else {
		fmt.Println("Some areas are still dirty.")
	}
}
