package main

import "fmt"

func main() {
	fmt.Println("2x2: Expecting 2")
	fmt.Println(robotPaths(2, 2))

	fmt.Println("2x3: Expecting 3")
	fmt.Println(robotPaths(2, 3))

	fmt.Println("3x3: Expecting 6")
	fmt.Println(robotPaths(3, 3))

	fmt.Println("4x4: Expecting 20")
	fmt.Println(robotPaths(4, 4))
}

func robotPaths(n, m int) int {
	return robotPathsIter(n, m, 0, 0)
}

func robotPathsIter(n, m, x, y int) int {
	// are we at the goal
	if x == n-1 && y == m-1 {
		return 1
	}

	var count int
	// move right
	if x != n-1 {
		count += robotPathsIter(n, m, x+1, y)
	}
	// move down
	if y != m-1 {
		count += robotPathsIter(n, m, x, y+1)
	}

	return count
}
