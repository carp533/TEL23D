package main

import (
	"fmt"
	"math"
)

// Es gibt mehrere Spielsteine in einem gemeinsamen Stapel.
// Zwei Spieler sind abwechselnd an der Reihe.
// Wenn ein Spieler an der Reihe ist, entfernt er einen, zwei oder drei Spielsteine vom Stapel.
// Der Spieler, der den letzten Stein nimmt, verliert das Spiel.
// 5 -> 4 -> 1
// 6 -> 5 ->

// "Spielfeld" -> Anzahl der Stein (int)

// Ermittelt die gültigen Züge
func possible_new_states(state int) []int {
	states := make([]int, 0)
	return states
}

// Bewertung der Endzustände
func evaluate(state int, max_turn bool) (score float64, spielende bool) {
	//Spieler gewinnt  +1
	//Spieler verliert -1
	return 0, false
}

// ermittelt den besten Zug
func best_move(state int) (score float64, move int) {
	// ermitteln der möglichen Züge
	// for Zug in Züge berechne minimax(Zug)
	// nehme das Maximum
	score = math.Inf(-1) //-1
	return score, move
}

// der Minimax Algorithmus
func minimax(state int, max_turn bool) (score float64) {
	// Prüfe, ob Spielende?
	// Wenn ja ->
	// gebe Wert (+1 / -1) zurück
	// Wenn nein ->
	// ermitteln der möglichen Züge
	// for Zug in Züge berechne minimax(Zug)
	// wenn MAX -> nehme das Maximum,
	// wenn MIN -> nehme das Minimum
	return score
}

func main() {
	fmt.Printf("minimax(5,true) =  %v\n", minimax(5, true))
	fmt.Printf("minimax(6,true) =  %v\n", minimax(6, true))
	// for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
	// 	score, stapel := best_move(v)
	// 	fmt.Printf("Stapel: %v, bester Zug %v (%v)\n", v, stapel, score)
	// }
}
