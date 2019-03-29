package main

import (
	"fmt"
)

func fib(x int) int {
	if x < 0 {
		fmt.Printf("ERR: cannot compute fib of non-positive x: %v\n",x)
		return x
	}
	
	if x == 2 || x == 1 || x == 0 {
		return 1
	}

	return fib(x-1) + fib(x-2)
	
}

func fibTest(x int) {
	fmt.Printf("fib of x [%v] : %v\n", x, fib(x))
}

func main() {
	fmt.Println("starting fib tests")

	for _, v := range([]int{1,2,3,4,5,6,7,8,10, -5}) {
		fibTest(v)
	}
	
	fmt.Println("ending fib tests")
}

