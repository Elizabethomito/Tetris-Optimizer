# Tetris-Optimizer

A **Go program** that reads a file of tetromino shapes and arranges them to form the **smallest possible square**. Each tetromino is represented by a unique uppercase letter in the solution.

---

## Features

- Parses tetrominoes from a text file (`#` for blocks, `.` for empty spaces).  
- Validates connectivity and format of tetrominoes.  
- Uses **recursive backtracking** to place tetrominoes efficiently.  
- Prints the solution board with tetrominoes labeled `A, B, C…`.  
- Handles invalid input gracefully by printing `ERROR`.  

---

## Requirements

- Go 1.24+  
- Standard Go packages only  

---

## Folder Structure
Tetris-Optimizer/
├── main.go # Entry point
├── parser/
│ ├── parser.go # Parsing tetrominoes
│ └── parser_test.go # Parser tests
├── solver/
│ ├── solver.go # Recursive solver
│ └── solver_test.go # Solver tests
├── model/
│ └── tetromino.go # Tetromino and Point structs
├── utils/
│ └── board.go # Board creation & printing
├── go.mod
├── go.sum
└── sample.txt # Sample tetromino input


---

## Input Format

- A tetromino is **4×4 characters**, `#` for block, `.` for empty.  
- Tetrominoes are separated by a **single empty line**.  

Example `sample.txt`:

...#
...#
...#
...#

....
..##
..##
....


---

## How to Run

1. Clone the repository:
```
git clone https://github.com/yourusername/Tetris-Optimizer.git
cd Tetris-Optimizer
```
2. Run the program with a text file:

```
   go run main.go sample.txt
```
   Example output:

```
ABBB.
ACCC.
A..C.
.....
   ```
Running Tests

Run all tests:
```
go test ./...
```

Run parser tests only:
```
go test ./parser -v
```
Run solver tests only:

```
go test ./solver -v
```
