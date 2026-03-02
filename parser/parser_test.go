package parser

import (
	"testing"
)

const validTetrominoes = `...#
...#
...#
...#

....
..##
..##
....`

const invalidTetrominoes = `...#
..#.
.#..
#...`

func TestParseFile_Valid(t *testing.T) {
	pieces, err := ParseFile(validTetrominoes)
	if err != nil {
		t.Fatalf("Expected valid tetrominoes, got error: %v", err)
	}

	if len(pieces) != 2 {
		t.Fatalf("Expected 2 tetrominoes, got %d", len(pieces))
	}

	if pieces[0].Letter != 'A' || pieces[1].Letter != 'B' {
		t.Fatalf("Expected letters A, B, got %c, %c", pieces[0].Letter, pieces[1].Letter)
	}
}

func TestParseFile_Invalid(t *testing.T) {
	_, err := ParseFile(invalidTetrominoes)
	if err == nil {
		t.Fatal("Expected error for invalid tetromino, got nil")
	}
}
