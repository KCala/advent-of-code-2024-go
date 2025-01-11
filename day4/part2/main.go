package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
)

type Board [][]rune

func (b Board) print() {
	for _, row := range b {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println("")
	}
}

func main() {
	input := file.MustReadInputAsLines()
	muls := parse(input)
	solution := solve(muls)
	fmt.Println(solution)
}

func parse(input []string) Board {
	rowsN := len(input)
	columnsN := len(input[0])

	var board Board
	board = make([][]rune, rowsN)
	for i := 0; i < len(board); i++ {
		board[i] = make([]rune, columnsN)
	}

	for l, line := range input {
		for c, char := range line {
			board[l][c] = char
		}
	}

	board.print()
	return board
}

func solve(input Board) int {

	total := 0

	for r, row := range input {
		for c := range row {
			total += countXmas(input, r, c)
		}
	}

	return total
}

func countXmas(board Board, row int, col int) int {
	if row <= 0 || row >= len(board)-1 || col <= 0 || col >= len(board[0])-1 {
		return 0
	}

	if board[row][col] != 'A' {
		return 0
	}

	backslashPresent := false // like this \
	slashPresent := false     // like this /

	if (board[row-1][col-1] == 'M' && board[row+1][col+1] == 'S') ||
		(board[row-1][col-1] == 'S' && board[row+1][col+1] == 'M') {
		backslashPresent = true
	}

	if (board[row+1][col-1] == 'M' && board[row-1][col+1] == 'S') ||
		(board[row+1][col-1] == 'S' && board[row-1][col+1] == 'M') {
		slashPresent = true
	}

	if backslashPresent && slashPresent {
		return 1
	} else {
		return 0
	}
}
