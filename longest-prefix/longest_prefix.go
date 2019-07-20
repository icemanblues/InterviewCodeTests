package main

import "fmt"

func longestPrefix(a []string) string {
	if a == nil || len(a) == 0 {
		return ""
	} else if len(a) == 1 {
		return a[0]
	}

	// 2 or more
	lp := cmp(a[0], a[1])
	for i := 2; i < len(a); i++ {
		lp = cmp(lp, a[i])
	}
	return lp
}

func cmp(a, b string) string {
	i := 0
	for ; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
	}
	return a[:i]
}

func main() {
	fmt.Println("Longest Prefix in Go")
	a := []string{}
	b := []string{"hello"}
	c := []string{"doodle", "double", "dave", "maritime"}
	d := []string{"floor", "florida", "fluid"}

	fmt.Printf("'' => %v\n", longestPrefix(a))
	fmt.Printf("hello => %v\n", longestPrefix(b))
	fmt.Printf("'' => %v\n", longestPrefix(c))
	fmt.Printf("fl => %v\n", longestPrefix(d))
}
