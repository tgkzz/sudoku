# Sudoku solver

The Sudoku Solver is a program that can solve any Sudoku puzzle. It's written in the Go language.

## Installation

If you have access to this repository, you can clone it with the following command:

```bash
git clone git@git.01.alem.school:zkhaypkh/sudoku.git
```

## Usage

You can pass your Sudoku problem through command line arguments when running the program. Use the following format: go run . **List-of-arguments**. The arguments should be strings representing the rows of the Sudoku puzzle.

Here is an example:

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```


## Algorithm

The algorithm used in our Sudoku Solver is based on a technique called "Backtracking". Here's a simplified explanation:

    Start at the first empty spot: The program starts at the first empty cell in the Sudoku grid.

    Try a number: It tries to put a number (from 1 to 9) in the current empty spot.

    Check if the number is valid: The program checks if this number is valid in the current row, column, and 3x3 box. If the number is valid, it moves on to the next empty spot and repeats the process.

    Backtrack if necessary: If the program finds that no numbers are valid in a particular spot, it "backtracks" to the previous empty spot and tries the next number. This process continues until it finds a number that works.

    Complete the grid: The program continues this process of trying numbers and backtracking until it has filled in all the empty spots in the grid with valid numbers, solving the Sudoku puzzle.


## Test

We have also provided tests for this exercise. To run them, use the following command:

```bash
go test -v
```