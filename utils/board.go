package utils

import "fmt"

// CreateBoard makes an empty square board
func CreateBoard(size int) [][]rune {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

// PrintBoard prints the board
func PrintBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}
