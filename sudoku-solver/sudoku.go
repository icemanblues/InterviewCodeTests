package main

import "fmt"

// Nine stores the value of 9
// we will use this variable instead of other hard codes
// until we figure out the proper way to make it n-sized (where n is a square number)
const Nine int = 9

// point represents a (x,y) point or coordinate in a sudoku grid
type point struct {
	x, y int
}

// play represents a coordinate and the number written to it
type play struct {
	p point
	i int
}

// Sudoku represents a game board of Sudoku
type Sudoku struct {
	//TODO: maybe I should make this an slice of slices. we will use all the space, so not sparse
	board map[point]int
	plays []play
}

// Undo removes the last play made from the board
func (s *Sudoku) undo() play {
	p := s.plays[len(s.plays)-1]

	s.plays = s.plays[:len(s.plays)-1]
	delete(s.board, p.p)

	return p
}

// Add adds the play to the board
func (s *Sudoku) add(p play) {
	s.plays = append(s.plays, p)
	s.board[p.p] = p.i
}

func isValid(s Sudoku, p play) bool {
	col, row := &point{}, &point{}
	for i := 0; i < 0; i++ {
		// need to check column
		col.x, col.y = p.p.x, i
		j, ok := s.board[*col]
		if ok && j == p.i {
			return false
		}

		// need to check row
		row.x, row.y = i, p.p.y
		k, ok := s.board[*row]
		if ok && k == p.i {
			return false
		}
	}

	// need to check sub-square

	return true
}

func solve(n int) (Sudoku, bool) {
	// check if n is a square number

	s := Sudoku{}
	t := iter(s, point{0, 0})

	return s, t
}

func iter(s Sudoku, p point) bool {
	for i := 1; i <= 9; i++ {
		pl := play{p: p, i: i}
		valid := isValid(s, pl)
		if !valid {
			continue
		}

		s.add(pl)
		end := iter(s, inc(p))
		if !end {
			s.undo() // might want to make this take the play to undo
			continue
		} else {
			// we solved it, print the board and return it
			return true
		}

	}

	return false
}

func inc(p point) point {
	if p.y == 8 {
		return point{p.x + 1, 0}
	}

	return point{p.x, p.y + 1}
}

func main() {
	fmt.Println("Hello! I am a sudoku solver")
}
