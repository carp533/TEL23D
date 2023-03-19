package hanoi

import "fmt"

// moveDisk moves a disk from from to to
func moveDisk(disk int, from string, to string) {
	fmt.Printf("Move disk %d from %s to %s\n", disk, from, to)
}

// solveTowersOfHanoi recursively solves the Towers of Hanoi problem
func SolveTowersOfHanoi(numDisks int, source string, auxiliary string, destination string) {
	if numDisks > 0 {
		// Move n-1 disks from source to auxiliary
		SolveTowersOfHanoi(numDisks-1, source, destination, auxiliary)

		// Move the nth disk from source to destination
		moveDisk(numDisks, source, destination)

		// Move n-1 disks from auxiliary to destination
		SolveTowersOfHanoi(numDisks-1, auxiliary, source, destination)
	}
}
