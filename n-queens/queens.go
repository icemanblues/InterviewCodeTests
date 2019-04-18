package main

import "fmt"

func queens(n int) {
	q := make([]int, n, n)
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

	// try numbers 0-(n-1)
	for j := 0; j < len(q); j++ {
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
	// iterate over columns
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
	fmt.Println("hello n queens")
	queens(8)
}
