package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// play represents a coordinate (x,y) and the number written to it (v)
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
	if n <= 0 {
		return nil, fmt.Errorf("argument n {%v} must be positive non-zero number", n)
	}
	if _, valid := isSquare(n); !valid {
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

// NewSudokuFromFile reads a sudoku file and creates a new SuDoKu board
func NewSudokuFromFile(filename string) (*Sudoku, error) {
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

// Undo removes the last play made from the board
func (s *Sudoku) undo() (play, bool) {
	if len(s.plays) == 0 {
		return play{}, false
	}

	p := s.plays[len(s.plays)-1]
	s.plays = s.plays[:len(s.plays)-1]
	s.board[p.x][p.y] = 0
	return p, true

}

// Add adds the play to the board
func (s *Sudoku) add(p play) {
	s.plays = append(s.plays, p)
	s.board[p.x][p.y] = p.v
}

func (s *Sudoku) isValid(p play) bool {
	n := len(s.board)

	for i := 0; i < n; i++ {
		// check the row of this play
		if p.v == s.board[i][p.y] {
			return false
		}

		// check the column of this play
		if p.v == s.board[p.x][i] {
			return false
		}
	}

	// need to check sub-square
	l := int(math.Sqrt(float64(n)))
	row := l * (p.x / l)
	col := l * (p.y / l)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if p.v == s.board[row+i][col+j] {
				return false
			}
		}
	}

	return true
}

// true if the number is a square number, integer square root
func isSquare(n int) (int, bool) {
	f := float64(n)
	s := math.Sqrt(f)
	return int(s), s == math.Floor(s)
}

// Solve solves the given SuDoKu puzzle
func Solve(s *Sudoku) bool {
	return iter(s, 0, 0)
}

func iter(s *Sudoku, x, y int) bool {
	// if we made it to the end of the board, we're done
	// we iterate across x last. for each x, we do a y iteration as per inc()
	if x == len(s.board) {
		return true
	}

	// there already is a number here that we can't change
	if s.board[x][y] != 0 {
		xi, yi := inc(x, y, len(s.board))
		return iter(s, xi, yi)
	}

	for i := 1; i <= len(s.board); i++ {
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
		}

		// we solved it, print the board and return it
		return true
	}

	return false
}

func inc(x, y, n int) (int, int) {
	if y == n-1 {
		return x + 1, 0
	}

	return x, y + 1
}

func main() {
	fmt.Println("Hello! I am a sudoku solver")

	easy, err := NewSudokuFromFile("easy.txt")
	if err != nil {
		fmt.Printf("Unable to load puzzle easy: %v\n", err)
		return
	}

	fmt.Println("attempting to solve...")
	sudokuSolve := Solve(easy)
	fmt.Printf("%v : %v\n", sudokuSolve, easy)
}
