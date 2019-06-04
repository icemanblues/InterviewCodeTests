package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var acc []int
	return iter(candidates, target, acc, 0)
}

func iter(candidates []int, target int, acc []int, idx int) [][]int {
	var cs [][]int

	for i := idx; i < len(candidates); i++ {
		t := target - candidates[i]
		if t < 0 {
			continue
		}

		acc = append(acc, candidates[i])
		if t == 0 {
			b := make([]int, len(acc))
			copy(b, acc)
			cs = append(cs, b)
			continue
		} else {
			r := iter(candidates, t, acc, i)
			cs = append(cs, r...)
		}
		acc = acc[:len(acc)-1]
	}

	return cs
}

func main() {
	test1c := []int{2, 3, 6, 7}
	test1t := 7
	a1 := combinationSum(test1c, test1t)
	fmt.Printf("Test1: %v\n", a1)

	test2c := []int{2, 3, 5}
	test2t := 8
	a2 := combinationSum(test2c, test2t)
	fmt.Printf("Test2: %v\n", a2)

	test3c := []int{7, 3, 2}
	test3t := 18
	a3 := combinationSum(test3c, test3t)
	fmt.Printf("Test3: %v\n", a3)
}
