package main

import "fmt"

func queens() {
	q := make([]int, 8, 8)
	fmt.Printf("board initialized: %v\n", q)

	ok := queenIter(q, 0)
	if ok {
		fmt.Printf("Solution Found: %v\n", q)
	}

}

// queens we've placed on the board
// i the current column (index) we are looking to place
func queenIter(q []int, i int) bool {
	if i >= len(q) {
		return true
	}

	// try numbers 0-7
	for j := 0; j < 8; j++ {
		// need to check if j is valid
		if valid := checkQueen(q, i, j); !valid {
			continue
		}

		q[i] = j
		ok := queenIter(q, i+1)
		if ok {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkQueen(q []int, i, j int) bool {

	for k := 0; k < i; k++ {
		// check rows
		if q[k] == j {
			return false
		}

		// check diagonal
		if abs(i-k) == abs(j-q[k]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("hello")
	queens()
}
