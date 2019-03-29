package main

import "fmt"

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
	board map[point]int
	plays []play
}

func (s *Sudoku) undo() play {
	p := s.plays[len(s.plays)-1]

	s.plays = s.plays[:len(s.plays)-1]
	delete(s.board, p.p)

	return p
}

func (s *Sudoku) add(p play) error {

}

func isValid(s Sudoku, p play) bool {

}

func main() {
	fmt.Println("Hello! I am a sudoku solver")
}
