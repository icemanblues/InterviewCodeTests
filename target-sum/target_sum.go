package main

import "fmt"

func main() {
	fmt.Println("target sum")

	a := []int{1,2,5,6,7,9,11}
	t := 9

	i, j := targetSum(a, t)
	fmt.Printf("(1,4) => %v, %v\n", i, j) 
}

func targetSum(a []int, t int) (int, int) {
	d := make(map[int]int)
	for i, e := range a {
		d[e] = i

		target := t - a[i]
		if j, ok := d[target]; ok {
			return i, j
		}
	}

	return -1, -1
}
