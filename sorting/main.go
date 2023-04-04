package main

import (
	"bubble"
	"fmt"
	"gosort"
	"heap"
	"insertion"
	"math/rand"
	"merge"
	"quick"
	"selection"
	"sort"
	"time"
)

func main() {

	slice2 := generateSlice(20)
	slice := make([]int, 20)

	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- heap     --- ", heap.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- bubble   --- ", bubble.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- selection -- ", selection.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- insertion -- ", insertion.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- quick    --- ", quick.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	fmt.Println("--- merge    --- ", merge.Sort(slice))
	copy(slice, slice2)
	fmt.Println("\n--- Unsorted --- ", slice)
	sort.Ints(slice)
	fmt.Println("--- inbuild  --- ", slice)
	gosort.DemoGoSort()
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
