package main

import (
	"fmt"
	"testing"
)

func TestNewSudoku(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		input := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
		_, err := NewSudoku(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		input := []string{".96.4...1", "1...6.1.4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
		_, err := NewSudoku(input)
		if err == nil {
			t.Errorf("Expected an error for invalid input, but didn't get one")
		}
	})

	t.Run("incorrect number of arguments", func(t *testing.T) {
		input := []string{"1", "2", "3", "4"}
		_, err := NewSudoku(input)
		if err == nil {
			t.Errorf("Expected an error for invalid input, but didn't get one")
		}
	})

	t.Run("more than one solution", func(t *testing.T) {
		input := make([]string, N)
		for i := range input {
			input[i] = "........."
		}
		_, err := NewSudoku(input)
		if err == nil {
			t.Errorf("Expected an error for invalid input, but didn't get one")
		}
	})
}

func TestSudoku_Solve(t *testing.T) {
	input := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	s, _ := NewSudoku(input)

	got := s.Solve()

	if got != true {
		t.Errorf("Expected Sudoku to be solved, but got %v", got)
	}
}

func TestSudoku_isValid(t *testing.T) {
	input := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	s, _ := NewSudoku(input)
	s.Solve()

	testCases := []struct {
		name string
		num  int
		row  int
		col  int
		want bool
	}{
		{name: "invalid row", num: 1, row: 0, col: 0, want: false},
		{name: "invalid column", num: 1, row: 3, col: 3, want: false},
		{name: "invalid block", num: 1, row: 1, col: 1, want: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := s.isValid(tt.num, tt.row, tt.col)
			if got != tt.want {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
		})
	}
}

func TestSudoku_Print(t *testing.T) {
	input := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	s, _ := NewSudoku(input)
	s.Solve()

	expected := [N][N]int{
		{3, 9, 6, 2, 4, 5, 7, 8, 1},
		{1, 7, 8, 3, 6, 9, 5, 2, 4},
		{5, 2, 4, 8, 1, 7, 3, 9, 6},
		{2, 8, 7, 9, 5, 1, 6, 4, 3},
		{9, 3, 1, 4, 8, 6, 2, 7, 5},
		{4, 6, 5, 7, 2, 3, 9, 1, 8},
		{7, 1, 2, 6, 3, 8, 4, 5, 9},
		{6, 5, 9, 1, 7, 4, 8, 3, 2},
		{8, 4, 3, 5, 9, 2, 1, 6, 7},
	}

	got := fmt.Sprintf("%v", *s)

	if got != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected\n%v\nbut got\n%v\n", expected, got)
	}
}
