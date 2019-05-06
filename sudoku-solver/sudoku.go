package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// play represents a coordinate and the number written to it
type play struct {
	x, y, v int
}

// Sudoku represents a game board of Sudoku
type Sudoku struct {
	board [][]int
	plays []play
}

// NewSudoku constructs a new Sudoku board in the empty state
func NewSudoku(n int) (*Sudoku, error) {
	if !isSquare(n) {
		return nil, fmt.Errorf("argument n {%v} is not a square number", n)
	}

	b := make([][]int, n, n)
	for i := range b {
		b[i] = make([]int, n, n)
	}
	return &Sudoku{
		board: b,
	}, nil
}

// Undo removes the last play made from the board
func (s *Sudoku) undo() play {
	p := s.plays[len(s.plays)-1]

	s.plays = s.plays[:len(s.plays)-1]
	s.board[p.x][p.y] = 0

	return p
}

// Add adds the play to the board
func (s *Sudoku) add(p play) {
	s.plays = append(s.plays, p)
	s.board[p.x][p.y] = p.v
}

func (s *Sudoku) isValid(p play) bool {
	n := len(s.board)
	// check the row of this play
	for i := 0; i < n; i++ {
		if p.v == s.board[i][p.y] {
			return false
		}
	}

	// check the column of this play
	for j := 0; j < n; j++ {
		if p.v == s.board[p.x][j] {
			return false
		}
	}

	// need to check sub-square
	// TODO: use int division to determine the "quadrant"
	// p.x / n
	// p.y / n

	return true
}

// true if the number is a square number, integer square root
func isSquare(n int) bool {
	f := float64(n)
	s := math.Sqrt(f)
	return s == math.Floor(s)
}

// Solve solves the given SuDoKu puzzle
func Solve(s *Sudoku) bool {
	return iter(s, 0, 0)
}

func iter(s *Sudoku, x, y int) bool {
	for i := 1; i <= 9; i++ {
		pl := play{x, y, i}
		valid := s.isValid(pl)
		if !valid {
			continue
		}

		s.add(pl)
		xi, yi := inc(x, y, len(s.board))
		end := iter(s, xi, yi)
		if !end {
			s.undo()
			continue
		} else {
			// we solved it, print the board and return it
			return true
		}
	}

	return false
}

func inc(x, y, n int) (int, int) {
	if y == n {
		return x + 1, 0
	}

	return x, y + 1
}

// LoadPuzzle reads a sudoku file and creates a new SuDoKu board
func LoadPuzzle(filename string) (*Sudoku, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var puzzle [][]int
	p := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, r := range line {
			if r != ' ' {
				row = append(row, int(r-'0'))
			} else {
				row = append(row, 0)
			}
		}
		if row != nil || len(row) != 0 {
			puzzle = append(puzzle, row)
		}
		p++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Sudoku{
		board: puzzle,
	}, nil
}

func main() {
	fmt.Println("Hello! I am a sudoku solver")

	fmt.Printf("is %v a square number? %v\n", 9, isSquare(9))
	fmt.Printf("is %v a square number? %v\n", 21, isSquare(21))

	empty, err := NewSudoku(9)
	if err != nil {
		fmt.Printf("Unable to create an empty puzzle: %v\n", err)
		return
	}
	fmt.Println(empty)

	// this one is supposed to fail
	fail, err := NewSudoku(21)
	if err != nil {
		fmt.Printf("Unable to create a new puzzle: %v\n", err)
	} else {
		fmt.Println(fail)
	}

	zero, err := LoadPuzzle("zero-sudoku.txt")
	if err != nil {
		fmt.Printf("Unable to load puzzle zero: %v\n", err)
		return
	}
	fmt.Println(zero)

	sudoku, err := LoadPuzzle("sudoku.txt")
	if err != nil {
		fmt.Printf("Unable to load puzzle zero: %v\n", err)
		return
	}
	fmt.Println(sudoku)
}
