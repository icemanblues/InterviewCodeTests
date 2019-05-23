package main

import "fmt"

// Reoccur detects the first re-occurring rune in a string
func Reoccur(s string) (rune, bool) {
	set := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := set[r]; ok {
			return r, true
		}
		set[r] = struct{}{}
	}

	return 'a', false
}

// ReoccurNoMem detects the first reoccurring character in a string using no memory
func ReoccurNoMem(s string) (rune, bool) {
	if len(s) == 0 {
		return 'z', false
	}

	runes := []rune(s)
	var last int
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				if last == 0 || j < last {
					last = j
				}
			}
		}
	}

	return runes[last], last != 0
}

func main() {
	fmt.Println("Reoccurring runes in a string")
}
