package main

import "hanoi"

func main() {
	numDisks := 3
	source := "A"
	auxiliary := "B"
	destination := "C"

	hanoi.SolveTowersOfHanoi(numDisks, source, auxiliary, destination)
}
