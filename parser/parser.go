package parser

import (
	"errors"
	"strings"
	"tetris_optimizer/model"
)

// ----------------------
// ParseTetromino
// ----------------------
func ParseTetromino(lines []string, letter rune) (model.Tetromino, error) {
	if len(lines) != 4 {
		return model.Tetromino{}, errors.New("wrong number of lines")
	}

	var points []model.Point
	for y, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) != 4 {
			return model.Tetromino{}, errors.New("line not 4 characters")
		}
		for x, ch := range line {
			if ch == '#' {
				points = append(points, model.Point{X: x, Y: y})
			} else if ch != '.' {
				return model.Tetromino{}, errors.New("invalid character")
			}
		}
	}

	if len(points) != 4 {
		return model.Tetromino{}, errors.New("must have 4 blocks")
	}

	// Check connectivity
	if !isConnected(points) {
		return model.Tetromino{}, errors.New("tetromino not connected")
	}

	// normalize to top-left
	minX, minY := 3, 3
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
	}
	for i := range points {
		points[i].X -= minX
		points[i].Y -= minY
	}

	return model.Tetromino{Blocks: points, Letter: letter}, nil
}

// ----------------------
// ParseFile
// ----------------------
func ParseFile(content string) ([]model.Tetromino, error) {
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	var pieces []model.Tetromino
	var blockLines []string
	letter := 'A'

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(blockLines) > 0 {
				tet, err := ParseTetromino(blockLines, letter)
				if err != nil {
					return nil, err
				}
				pieces = append(pieces, tet)
				letter++
				blockLines = []string{}
			}
			continue
		}

		blockLines = append(blockLines, line)

		if i == len(lines)-1 && len(blockLines) > 0 {
			tet, err := ParseTetromino(blockLines, letter)
			if err != nil {
				return nil, err
			}
			pieces = append(pieces, tet)
			letter++
			blockLines = []string{}
		}
	}

	if len(pieces) == 0 {
		return nil, errors.New("no tetrominoes found")
	}

	return pieces, nil
}

// ----------------------
// Connectivity check
// ----------------------
func isConnected(points []model.Point) bool {
	if len(points) != 4 {
		return false
	}

	visited := make(map[model.Point]bool)
	var dfs func(p model.Point)
	dfs = func(p model.Point) {
		if visited[p] {
			return
		}
		visited[p] = true
		for _, n := range neighbors(p) {
			if contains(points, n) {
				dfs(n)
			}
		}
	}
	dfs(points[0])
	return len(visited) == 4
}

func neighbors(p model.Point) []model.Point {
	return []model.Point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
}

func contains(points []model.Point, p model.Point) bool {
	for _, pt := range points {
		if pt == p {
			return true
		}
	}
	return false
}
