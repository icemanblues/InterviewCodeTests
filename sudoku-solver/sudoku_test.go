package main

import (
	"testing"
)

func TestIsSquare(t *testing.T) {
	tests := []struct {
		name      string
		arg, root int
		expected  bool
	}{
		{"4 is square", 4, 2, true},
		{"16 is square", 16, 4, true},
		{"not square", 21, 2, false},
		{"negative square", -81, -9, false},
		{"negative non-square", -21, -9, false},
	}

	for _, tt := range tests {
		r, e := isSquare(tt.arg)
		if e != tt.expected {
			t.Fatalf("%v Expected boolean: %v actual: %v", tt.name, tt.expected, e)
		}
		// only check root if expected=true (it is a square number)
		if tt.expected && r != tt.root {
			t.Fatalf("%v Expected root: %v but actual: %v", tt.name, tt.root, r)
		}
	}
}

func TestInc(t *testing.T) {
	tests := []struct {
		x, y, n int
		ex, ey  int
	}{
		{0, 0, 9, 0, 1},
		{0, 8, 9, 1, 0},
		{2, 5, 9, 2, 6},
		{8, 1, 9, 8, 2},
		{8, 8, 9, 9, 0},
	}

	for _, tt := range tests {
		ax, ay := inc(tt.x, tt.y, tt.n)
		if ax != tt.ex {
			t.Fatalf("inc{%d, %d, %d) expected x: %d actual: %d", tt.x, tt.y, tt.n, tt.ex, ax)
		}
		if ay != tt.ey {
			t.Fatalf("inc{%d, %d, %d) expected y: %d actual: %d", tt.x, tt.y, tt.n, tt.ey, ay)
		}
	}
}

func TestNewSudoku(t *testing.T) {
	tests := []struct {
		n   int
		err bool
	}{
		{1, false}, {9, false}, {21, true}, {0, true}, {-81, true},
	}

	for _, tt := range tests {
		s, err := NewSudoku(tt.n)
		if tt.err != (err != nil) {
			t.Fatalf("NewSudoku(%d) Expecting existence of error: %v Actual: %v", tt.n, tt.err, err)
		}

		if !tt.err {
			if l := len(s.plays); l != 0 {
				t.Fatalf("NewSudoku(%d) A newly constructed sudoku should have no plays in it. Found: %v %v", tt.n, l, s.plays)
			}

			if l := len(s.board); l != tt.n {
				t.Fatalf("NewSudoku(%d) Incorrect number of rows. Expecting %v Actual %v", tt.n, tt.n, l)
			}

			for i := 0; i < tt.n; i++ {
				if l := len(s.board[i]); l != tt.n {
					t.Fatalf("NewSudoku(%d) Incorrect number of cols on row %v. Expecting %v Actual %v", tt.n, i, tt.n, l)
				}
			}

			for i := 0; i < tt.n; i++ {
				for j := 0; j < tt.n; j++ {
					if k := s.board[i][j]; k != 0 {
						t.Fatalf("NewSudoku(%d) Non zero value (%d) found on row %v col %v", tt.n, k, i, j)
					}
				}
			}
		}
	}
}

func TestNewSudokuFromFile(t *testing.T) {
	tests := []struct {
		file     string
		num, sum int
	}{
		{"easy.txt", 34, 168},
		{"zero-sudoku.txt", 0, 0},
	}

	for _, tt := range tests {
		s, err := NewSudokuFromFile(tt.file)
		if err != nil {
			t.Fatal(err)
		}

		if l := len(s.plays); l != 0 {
			t.Fatalf("Sudoku from file has plays. Expected 0, Actual %d", l)
		}

		count, sigma := 0, 0
		for i := 0; i < len(s.board); i++ {
			for j := 0; j < len(s.board[i]); j++ {
				if s.board[i][j] != 0 {
					count++
					sigma += s.board[i][j]
				}
			}
		}
		if count != tt.num {
			t.Fatalf("Expecting count of non-zero values: %d Actual: %d", tt.num, count)
		}
		if sigma != tt.sum {
			t.Fatalf("Expecting summation of all values: %d Actual: %d", tt.sum, sigma)
		}
	}
}

func TestAdd(t *testing.T) {
	s, err := NewSudoku(9)
	if err != nil {
		t.Fatalf("TestAdd: Unable to create an empty Sudoku: %v", err)
	}

	if len(s.plays) != 0 {
		t.Fatalf("starting Sudoku is not empty")
	}

	p1 := play{1, 2, 3}
	p2 := play{4, 5, 6}

	s.add(p1)
	if l := len(s.plays); l != 1 {
		t.Fatalf("first add is not correct. Expecting len(plays) 1 but actual: %d", l)
	}
	if s.plays[0] != p1 {
		t.Fatalf("The play we added is not there. Expected: %v Actual: %v", s.plays[0], p1)
	}

	s.add(p2)
	if l := len(s.plays); l != 2 {
		t.Fatalf("second add is not correct. Expecting len(plays) 2 but actual: %d", l)
	}
	if s.plays[1] != p2 {
		t.Fatalf("The play we added is not there. Expected: %v Actual: %v", s.plays[1], p2)
	}
}

func TestUndo(t *testing.T) {
	s, err := NewSudoku(4)
	if err != nil {
		t.Fatalf("Unable to create an empty sudoku")
	}

	_, ok := s.undo()
	if ok {
		t.Fatalf("Undo on empty sudoku should return false: %v", ok)
	}

	p := play{1, 2, 3}
	s.add(p)
	actual, ok := s.undo()
	if !ok {
		t.Fatalf("Undo on non-empty sudoku is valid. should return true. %v", ok)
	}
	if actual != p {
		t.Fatalf("s.undo Expected: %v Actual: %v", p, actual)
	}
}

func TestIsValid(t *testing.T) {
	s, err := NewSudokuFromFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		p     play
		valid bool
	}{
		{play{0, 2, 3}, true},
		{play{0, 2, 4}, false},

		{play{2, 0, 2}, true},
		{play{2, 0, 1}, false},

		{play{1, 1, 3}, true},
		{play{1, 1, 1}, false},
	}
	for _, tt := range tests {
		if actual := s.isValid(tt.p); actual != tt.valid {
			t.Fatalf("isValid(%v) Expected: %v Actual %v", tt.p, tt.valid, actual)
		}
	}
}
