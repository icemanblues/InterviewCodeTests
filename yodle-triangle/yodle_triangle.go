package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readTriangle(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var triangle [][]int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		ints := make([]int, len(nums), len(nums))
		for i, n := range nums {
			a, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			ints[i] = a
		}
		triangle = append(triangle, ints)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return triangle, nil
}

func topDown(triangle [][]int) int {
	return topDownRec(triangle, 0, 0)
}

func topDownRec(triangle [][]int, r, c int) int {
	if r >= len(triangle) {
		return 0
	}
	return triangle[r][c] + max(topDownRec(triangle, r+1, c), topDownRec(triangle, r+1, c+1))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func bottomUp(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		top := triangle[i]
		bot := triangle[i+1]
		for j := range top {
			top[j] += max(bot[j], bot[j+1])
		}
	}

	return triangle[0][0]
}

func main() {
	fmt.Println("Hello World")

	sample, err := readTriangle("sample.txt")
	if err != nil {
		fmt.Printf("Unable to continue: %v", err)
		return
	}
	fmt.Printf("Sample: %v\n", sample)

	fmt.Printf("Sample solved with top down: %v\n", topDown(sample))
	fmt.Printf("Sample solved with bottom up: %v\n", bottomUp(sample))

	hard, err := readTriangle("triangle.txt")
	if err != nil {
		fmt.Printf("Unable to proceed: %v", err)
		return
	}
	fmt.Printf("HARD solved with bottom up: %v\n", bottomUp(hard))
}
