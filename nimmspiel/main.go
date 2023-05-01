package main

import (
	"fmt"
)

// Es gibt mehrere Spielsteine in einem gemeinsamen Stapel.
// Zwei Spieler sind abwechselnd an der Reihe.
// Wenn ein Spieler an der Reihe ist, entfernt er einen, zwei oder drei Spielsteine vom Stapel.
// Der Spieler, der den letzten Stein nimmt, verliert das Spiel.

// Ermittelt die gültigen Züge
func possible_new_states(state int) []int {
	//TODO
	return nil
}

// Bewertung der Endzustände
func evaluate(state int, max_turn bool) (score float64, spielende bool) {
	//TODO
	return 0, false
}

// ermittelt den besten Zug
func best_move(state int) (score float64, move int) {
	//TODO
	return 0, 0
}

// der Minimax Algorithmus
func minimax(state int, max_turn bool) float64 {
	//TODO
	return 0
}

func main() {
	fmt.Printf("minimax(5,true) =  %v\n", minimax(5, true))
	fmt.Printf("minimax(6,true) =  %v\n", minimax(6, true))
	// for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
	// 	score, stapel := best_move(v)
	// 	fmt.Printf("Stapel: %v, bester Zug %v (%v)\n", v, stapel, score)
	// }
}
