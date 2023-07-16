package main

import (
	"fmt"
	"os"
)

const (
	N     = 9
	Empty = 0
)

type Sudoku [N][N]int

func NewSudoku(input []string) (*Sudoku, error) {
	var s Sudoku

	count := 0

	seenRows := make([]map[int]bool, N) // slice of maps to keep track of the numbers seen in each row & col
	seenCols := make([]map[int]bool, N)

	for i := 0; i < N; i++ {
		seenRows[i] = make(map[int]bool) // empty map for the current row to track the numbers seen
		seenCols[i] = make(map[int]bool)
	}

	for i, row := range input {
		if len(row) != N {
			return nil, fmt.Errorf("Invalid input: incorrect number of columns in row %d", i+1)
		}
		for j, c := range row {
			if c == '.' {
				s[i][j] = Empty // if dot, it is an empty cell
				count++
				if count > 64 {
					return nil, fmt.Errorf("This puzzle has more than 1 possible solution")
				}
			} else {
				s[i][j] = int(c - '0') // if not a dot, it is a digit
				if s[i][j] < 1 || s[i][j] > 9 {
					return nil, fmt.Errorf("Invalid input: invalid digit '%c' at position (%d, %d)", c, i+1, j+1)
				}

				if seenRows[i][s[i][j]] {
					return nil, fmt.Errorf("Invalid input: repeating value '%d' in row %d", s[i][j], i+1)
				}
				if seenCols[j][s[i][j]] {
					return nil, fmt.Errorf("Invalid input: repeating value '%d' in column %d", s[i][j], j+1)
				}

				seenRows[i][s[i][j]] = true
				seenCols[j][s[i][j]] = true
			}
		}
	}
	return &s, nil // return created sudoku object and nil error if input is valid
}

func (s *Sudoku) Solve() bool {
	for row := 0; row < N; row++ { // iterate through rows and cols
		for col := 0; col < N; col++ {
			if s[row][col] == Empty { // is cell empty, fill it with digit
				for num := 1; num <= 9; num++ {
					if s.isValid(num, row, col) { // try digits from 1 to 9 and test if it's valid
						s[row][col] = num
						if s.Solve() { // recursive call
							return true
						}
						s[row][col] = Empty // if the current digit not valid, try another
					}
				}
				return false // if no digits are valid, backtrack
			}
		}
	}
	return true // if all cells filled, puzzle is solved
}

func (s *Sudoku) isValid(num, row, col int) bool {
	for x := 0; x < 9; x++ {
		if s[row][x] == num { // check if current row contains testing value
			return false
		}
	}

	for x := 0; x < 9; x++ {
		if s[x][col] == num { // check if current column contains testing value
			return false
		}
	}

	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i+startRow][j+startCol] == num { // check if current subgrid contains testing value
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) Print() {
	for _, row := range s {
		for _, val := range row {
			fmt.Printf("%d ", val) // print solving row by row
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != N+1 {
		fmt.Println("Error: Invalid number of arguments")
		os.Exit(1)
	}

	s, err := NewSudoku(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if !s.Solve() {
		fmt.Println("Error: Unable to solve the Sudoku puzzle")
		os.Exit(1)
	}

	s.Print()
}
