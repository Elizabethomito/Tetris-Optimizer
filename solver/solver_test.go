package solver

import (
	"testing"
	"tetris_optimizer/model"
	"tetris_optimizer/utils"
)

func TestSolver_Simple(t *testing.T) {
	pieces := []model.Tetromino{
		{
			Blocks: []model.Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			Letter: 'A',
		},
		{
			Blocks: []model.Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			Letter: 'B',
		},
	}

	board := utils.CreateBoard(4)

	if !Solve(board, pieces, 0) {
		t.Fatal("Expected solver to find a solution")
	}

	// Check letters are placed
	foundA, foundB := false, false
	for _, row := range board {
		for _, cell := range row {
			if cell == 'A' {
				foundA = true
			}
			if cell == 'B' {
				foundB = true
			}
		}
	}

	if !foundA || !foundB {
		t.Fatal("Solver did not place all pieces correctly")
	}
}
