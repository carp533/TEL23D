package main

import "fmt"

/*Eine Liste stellt ein Histogram dar. Es soll die Fläche des größten
Rechtecks berechnet werden. die Balken des Histogramms haben Breite 1 und
Höhe ist der Wert aus der Liste.
Die Lösung erfolgt mit Hilfe von zwei Slices, in denen jeweils der Index
gespeichert wird, wie weit das Rechteck nach links, bzw. rechts gehen kann.
Falls das Rechteck bis ganz links/rechts gehen kann, ist der Index -1 bzw. len(heights) */
func minl(heights []int) []int {
	lessFromLeft := make([]int, len(heights))
	lessFromLeft[0] = -1
	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}
	return lessFromLeft
}
func minr(heights []int) []int {
	lessFromRight := make([]int, len(heights))
	lessFromRight[len(heights)-1] = len(heights)
	//TODO
	return nil
}
func largestRectangleArea(heights []int) int {
	lessFromLeft := minl(heights)
	lessFromRight := minr(heights)
	fmt.Println(lessFromLeft)
	fmt.Println(lessFromRight)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		//TODO: wie wird die Fläche berechnet? warum wird hier -1 genommen?
		maxArea = max(maxArea, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}
	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func ExampleRect() {
	x := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(x)
	fmt.Println(largestRectangleArea(x))
	//Output:
	//10
}
