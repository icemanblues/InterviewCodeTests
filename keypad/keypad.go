package main

import "fmt"

var keypad = map[rune][]rune{
	'0': []rune{'0'},
	'1': []rune{'1'},
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func t9(number string) [][]rune {
	acc := make([][]rune, 0)
	runes := []rune(number)
	t9iter(runes, 0, nil, &acc)
	return acc
}

func t9iter(number []rune, idx int, l []rune, acc *[][]rune) {
	if idx >= len(number) {
		ll := append([]rune(nil), l...)
		*acc = append(*acc, ll)
	} else {
		possibles := keypad[number[idx]]
		for _, p := range possibles {
			l = append(l, p)
			t9iter(number, idx+1, l, acc)
			l = l[:len(l)-1]
		}
	}
}

func main() {
	fmt.Println("keypad")
	fmt.Printf("23 => %c\n", t9("23"))
}
