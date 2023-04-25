package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var cnt int

const (
	SIZE       = 3
	PLAYER_MAX = "X"
	PLAYER_MIN = "O"
	EMPTY      = " "
)

type board struct {
	board [][]string
	turn  string
}

func NewBoard() *board {
	b := board{}
	b.board = [][]string{{EMPTY, EMPTY, EMPTY}, {EMPTY, EMPTY, EMPTY}, {EMPTY, EMPTY, EMPTY}}
	b.turn = PLAYER_MAX
	return &b
}

func getPossibleMoves(b board) ([][][]string, [][]int) {
	player := b.turn
	var moves [][][]string
	var moveIndex [][]int
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			boardcopy := make([][]string, SIZE)
			for k := 0; k < SIZE; k++ {
				boardcopy[k] = make([]string, SIZE)
				copy(boardcopy[k], b.board[k])
			}
			if boardcopy[i][j] == EMPTY {
				moveIndex = append(moveIndex, []int{i, j})
				boardcopy[i][j] = player
				moves = append(moves, boardcopy)
			}
		}
	}
	return moves, moveIndex
}

func minimax(b board) (float64, []int) {
	cnt++
	winner := checkWin(b)
	if winner != "" {
		if winner == PLAYER_MAX {
			return 1, nil
		} else if winner == PLAYER_MIN {
			return -1, nil
		} else if winner == EMPTY {
			return 0, nil
		}
	}

	possibleMoves, moveIndex := getPossibleMoves(b)
	bestMove := make([]int, 0)
	if b.turn == PLAYER_MAX {
		maxEval := math.Inf(-1)
		for index, item := range possibleMoves {
			tempBoard := board{item, PLAYER_MIN}
			eval, move := minimax(tempBoard)
			if eval >= maxEval {
				maxEval = eval
				bestMove = moveIndex[index]
				if move == nil {
					break
				}
			}
		}
		return maxEval, bestMove
	} else {
		minEval := math.Inf(1)
		for index, item := range possibleMoves {
			tempBoard := board{item, PLAYER_MAX}
			eval, move := minimax(tempBoard)
			if eval <= minEval {
				minEval = eval
				bestMove = moveIndex[index]
				if move == nil {
					break
				}
			}
		}
		return minEval, bestMove
	}

}

func checkWin(b board) string {
	for i := 0; i < SIZE; i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] != EMPTY {
			return b.board[i][0]
		}
		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] != EMPTY {
			return b.board[0][i]
		}
	}
	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[1][1] != EMPTY {
		return b.board[0][0]
	}
	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[1][1] != EMPTY {
		return b.board[0][2]
	}
	isDraw := true
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b.board[i][j] == EMPTY {
				isDraw = false
			}
		}
	}
	if isDraw {
		return EMPTY
	}
	return ""
}

func PlayGame() {
	b := NewBoard()
	//possibleMoves, _ := getPossibleMoves(*b)
	//fmt.Println(possibleMoves)

	opponent := choose_opponent()
	var xInp, yInp int
	for {
		printBoard(b.board)
		if opponent == 1 || (b.turn == PLAYER_MAX && opponent == 2) || (b.turn == PLAYER_MIN && opponent == 3) {
			xInp, yInp = getInput(*b)
		} else {
			_, move := minimax(*b)
			xInp = move[0]
			yInp = move[1]
		}
		b.board[xInp][yInp] = b.turn
		result := checkWin(*b)
		if result == EMPTY {
			printBoard(b.board)
			fmt.Println("Unentschieden.")
			return
		}
		if result != "" {
			printBoard(b.board)
			fmt.Printf("%s hat das Spiel gewonnen!", b.turn)
			return
		}
		if b.turn == PLAYER_MAX {
			b.turn = PLAYER_MIN
		} else {
			b.turn = PLAYER_MAX
		}
	}
}

func getInput(b board) (int, int) {
	reader := bufio.NewReader(os.Stdin)
	var X, Y int
	for {
		fmt.Print("Eingabe Zeile (0,1,2):")
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
		fmt.Print("Eingabe Spalte (0,1,2):")
		y, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Da hat was nicht geklappt. Versuche es nochmal.")
			continue
		}
		y = strings.Replace(y, "\n", "", -1)
		y = strings.Replace(y, "\r", "", -1)
		Y, err = strconv.Atoi(y)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Da hat was nicht geklappt. Versuche es nochmal.")
			continue
		}
		if X > 2 || Y > 2 || X < 0 || Y < 0 || b.board[X][Y] != EMPTY {
			fmt.Println("Ungültige Eingabe. Versuche es nochmal.")
			continue
		}
		break
	}
	return X, Y
}

func printBoard(board [][]string) {
	fmt.Println("---------")
	for ix := 0; ix < SIZE; ix++ {
		for iy := 0; iy < SIZE; iy++ {
			fmt.Print("|", board[ix][iy], "|")
		}
		fmt.Println("\n---------")
	}
}

func choose_opponent() int {
	return 2
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Println("Do you want to play against a human or the computer?")
	// 	fmt.Println("1: Human")
	// 	fmt.Println("2: Computer")
	// 	choice, err := reader.ReadString('\n')
	// 	choice = strings.Replace(choice, "\n", "", -1)
	// 	choice = strings.Replace(choice, "\r", "", -1)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		fmt.Println("Error while getting input. Try again.")
	// 		continue
	// 	}
	// 	if choice == "1" || choice == "2" {
	// 		intChoice, err := strconv.Atoi(choice)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			fmt.Println("Error while getting input. Try again.")
	// 			continue
	// 		}
	// 		if intChoice == 2 {
	// 			fmt.Println("Who goes first?")
	// 			fmt.Println("1: You")
	// 			fmt.Println("2: Computer")
	// 			firstPlayer, err := reader.ReadString('\n')
	// 			if err != nil {
	// 				fmt.Println(err)
	// 				fmt.Println("Error while getting input. Try again.")
	// 				continue
	// 			}
	// 			firstPlayer = strings.Replace(firstPlayer, "\n", "", -1)
	// 			firstPlayer = strings.Replace(firstPlayer, "\r", "", -1)
	// 			if firstPlayer == "1" {
	// 				return 2
	// 			} else if firstPlayer == "2" {
	// 				return 3
	// 			} else {
	// 				fmt.Println("Invalid Input. Try again.")
	// 				continue
	// 			}
	// 		}
	// 		return intChoice
	// 	} else {
	// 		fmt.Println("Invalid Input. Try again.")
	// 	}
	// }
}

func main() {
	PlayGame()
	fmt.Println("Rekursionen", cnt)
}
