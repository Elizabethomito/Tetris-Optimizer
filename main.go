package main

import (
	"fmt"
	"os"
	"tetris_optimizer/parser"
	"tetris_optimizer/solver"
	"tetris_optimizer/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	file := os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces, err := parser.ParseFile(string(content))
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// minimum board size
	boardSize := 2
	for boardSize*boardSize < len(pieces)*4 {
		boardSize++
	}

	var board [][]rune

	for {
		board = utils.CreateBoard(boardSize)
		if solver.Solve(board, pieces, 0) {
			break
		}
		boardSize++
	}

	utils.PrintBoard(board)
}
