package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var field = [9][9]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}
	if field[y][x] != 0 {
		return solve(next(x, y))
	}
	for i := range [9]int{} {
		var v = i + 1
		if canPut(x, y, v) {
			field[y][x] = v
			if solve(next(x, y)) {
				return true
			}
			field[y][x] = 0
		}
	}
	return false

}

func main() {
	field = parseInput("53..7....6..195....98....6.8...6...34..8.3..17...2...6.6....28....419..5....8..79")
	printBoard(field)
	if solve(0, 0) {
		fmt.Println("Found solution")
		printBoard(field)
	} else {
		fmt.Println("No solution")
	}
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}

// func possible(y, x, n int, board *[9][9]int) bool {
// 	for i := 0; i < 9; i++ {
// 		if board[y][i] == n {
// 			return false
// 		}
// 		if board[i][x] == n {
// 			return false
// 		}
// 	}
// 	x0 := int((x / 3) * 3)
// 	y0 := int((y / 3) * 3)
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if board[y0+i][x0+i] == n {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func solve(board *[9][9]int) bool {
// 	for y := 0; y < 9; y++ {
// 		for x := 0; x < 9; x++ {
// 			if board[y][x] == 0 {
// 				for n := 1; n < 9+1; n++ {
// 					if possible(y, x, n, board) {
// 						board[y][x] = n
// 						if solve(board) {
// 							printBoard(*board)
// 							return true
// 						}
// 						board[y][x] = 0
// 					}
// 				}
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func main() {

// 	grid := parseInput("3..1.5.68......7.19.1....3...7.26...5.......3...87.4...3....8.51.5......79.4.1...")
// 	// grid = parseInput("53..7....6..195....98....6.8...6...34..8.3..17...2...6.6....28....419..5....8..79")
// 	fmt.Println(possible(1, 1, 2, &grid))
// 	printBoard(grid)
// 	if solve(&grid) {
// 		fmt.Println("The Sudoku was solved successfully:")
// 		printBoard(grid)
// 	} else {
// 		fmt.Printf("The Sudoku can't be solved.")
// 	}

// }
