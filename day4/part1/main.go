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

	return total;
}

func countXmas(board Board, row int, col int) int {
    count := 0
    // horizontal backwards
    if col >= 3 {
        if board[row][col] == 'X' && board[row][col-1] == 'M' && board[row][col-2] == 'A' && board[row][col-3] == 'S' {
            count++
        }
    }
    // horizontal forward
    if col <= len(board[0])-4 {
        if board[row][col] == 'X' && board[row][col+1] == 'M' && board[row][col+2] == 'A' && board[row][col+3] == 'S' {
            count++
        }
    }
    // vertical upwards
    if row >= 3 {
        if board[row][col] == 'X' && board[row-1][col] == 'M' && board[row-2][col] == 'A' && board[row-3][col] == 'S' {
            count++
        }
    }
    // vertical downwards
    if row <= len(board)-4 {
        if board[row][col] == 'X' && board[row+1][col] == 'M' && board[row+2][col] == 'A' && board[row+3][col] == 'S' {
            count++
        }
    }
    // diagonal upwards right
    if row >= 3 && col <= len(board[0])-4 {
        if board[row][col] == 'X' && board[row-1][col+1] == 'M' && board[row-2][col+2] == 'A' && board[row-3][col+3] == 'S' {
            count++
        }
    }
    // diagonal upwards left
    if row >= 3 && col >= 3 {
        if board[row][col] == 'X' && board[row-1][col-1] == 'M' && board[row-2][col-2] == 'A' && board[row-3][col-3] == 'S' {
            count++
        }
    }
    // diagonal downwards right
    if row <= len(board)-4 && col <= len(board[0])-4 {
        if board[row][col] == 'X' && board[row+1][col+1] == 'M' && board[row+2][col+2] == 'A' && board[row+3][col+3] == 'S' {
            count++
        }
    }
    // diagonal downwards left
    if row <= len(board)-4 && col >= 3 {
        if board[row][col] == 'X' && board[row+1][col-1] == 'M' && board[row+2][col-2] == 'A' && board[row+3][col-3] == 'S' {
            count++
        }
    }
    return count
}
