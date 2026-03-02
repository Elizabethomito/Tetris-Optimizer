package solver

import "tetris_optimizer/model"

// canPlace checks if tetromino fits at position (row, col)
func canPlace(board [][]rune, t model.Tetromino, row, col int) bool {
	size := len(board)
	for _, p := range t.Blocks {
		r := row + p.Y
		c := col + p.X
		if r < 0 || r >= size || c < 0 || c >= size {
			return false
		}
		if board[r][c] != '.' {
			return false
		}
	}
	return true
}

// place puts the tetromino on the board
func place(board [][]rune, t model.Tetromino, row, col int) {
	for _, p := range t.Blocks {
		board[row+p.Y][col+p.X] = t.Letter
	}
}

// remove removes the tetromino from the board
func remove(board [][]rune, t model.Tetromino, row, col int) {
	for _, p := range t.Blocks {
		board[row+p.Y][col+p.X] = '.'
	}
}

// Solve tries to place all tetrominoes recursively
func Solve(board [][]rune, pieces []model.Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}

	size := len(board)
	t := pieces[index]

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if canPlace(board, t, r, c) {
				place(board, t, r, c)
				if Solve(board, pieces, index+1) {
					return true
				}
				remove(board, t, r, c)
			}
		}
	}

	return false
}
