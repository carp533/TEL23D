package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	PLAYER_COMPUTER = true
	PLAYER_HUMAN    = false
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
	for i := 1; i < 4 && state-i >= 0; i++ {
		states = append(states, state-i)
	}
	return states
}

// Bewertung der Endzustände
func evaluate(state int, max_turn bool) (score float64, spielende bool) {
	//Spieler gewinnt  +1
	//Spieler verliert -1
	if state == 0 {
		if max_turn {
			return 1, true
		}
		return -1, true
	}
	return 0, false
}

// ermittelt den besten Zug
func best_move(state int) (score float64, move int) {
	// ermitteln der möglichen Züge
	// for Zug in Züge berechne minimax(Zug)
	// nehme das Maximum
	possible_new_states := possible_new_states(state)
	score = math.Inf(-1)
	for _, new_state := range possible_new_states {
		eval := minimax(new_state, false)
		if eval >= score {
			score = eval
			move = new_state
		}
	}
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
	score, end := evaluate(state, max_turn)
	if end == true {
		return float64(score)
	}
	possible_new_states := possible_new_states(state)
	if max_turn {
		maxEval := math.Inf(-1)
		for _, new_state := range possible_new_states {
			eval := minimax(new_state, false)
			if eval >= maxEval {
				maxEval = eval
			}
		}
		return maxEval
	}
	minEval := math.Inf(1)
	for _, new_state := range possible_new_states {
		eval := minimax(new_state, true)
		if eval <= minEval {
			minEval = eval
		}
	}
	return minEval
}

func main() {
	// getInput(2)
	PlayGame(7)
	// fmt.Printf("minimax(5,true) =  %v\n", minimax(5, true))
	// fmt.Printf("minimax(6,true) =  %v\n", minimax(6, true))
	// for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
	// 	score, stapel := best_move(v)
	// 	fmt.Printf("Stapel: %v, bester Zug %v (%v)\n", v, stapel, score)
	// }
}

func getInput(state int) (newstate int) {
	reader := bufio.NewReader(os.Stdin)
	var X int
	for {
		fmt.Printf("Steine %v. Mache einen Zug:", state)
		x, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Da hat was nicht geklappt. Versuche es nochmal.")
			continue
		}
		x = strings.Replace(x, "\n", "", -1)
		x = strings.Replace(x, "\r", "", -1)
		X, err = strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Da hat was nicht geklappt. Versuche es nochmal.")
			continue
		}
		if state-X < 0 {
			fmt.Println("Du hast dir zu viele Steine genommen!")
			continue
		}
		if X < 1 || X > 3 {
			fmt.Println("Du darfst nur 1-3 Steine nehmen!")
			continue
		}
		return state - X
	}
}

func PlayGame(initstate int) {
	player := PLAYER_HUMAN
	new_state := initstate
	for {
		if player == PLAYER_COMPUTER {
			player = PLAYER_HUMAN
			_, new_state = best_move(new_state)
			fmt.Printf("Computer moves %v\n", new_state)
		} else {
			new_state = getInput(new_state)
			player = PLAYER_COMPUTER
			fmt.Printf("Human moves %v\n", new_state)
		}
		_, ende := evaluate(new_state, true)
		if ende {
			if PLAYER_HUMAN {
				fmt.Println("Comuter wins")
			} else {
				fmt.Println("Human wins")
			}
			return
		}
	}
}
