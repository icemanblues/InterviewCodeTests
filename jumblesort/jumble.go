package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// JumbleSort sorts a line, so that the strings are in alphabetical order, and the numbers are in
// numeric order. And if the element was a string, it remains a string.
func JumbleSort(line string) []string {
	elements := strings.Split(line, " ")

	var s []string
	var n []int
	var b []bool // true means that position is an int, false is a string

	for _, e := range elements {
		// try to convert e to an int
		num, err := strconv.Atoi(e)
		if err != nil {
			s = append(s, e)
			b = append(b, false)
		} else {
			n = append(n, num)
			b = append(b, true)
		}
	}

	// sort the partitions
	sort.Ints(n)
	sort.Strings(s)

	// merge it back together
	var jumble []string
	for _, pos := range b {
		if pos {
			jumble = append(jumble, strconv.Itoa(n[0]))
			n = n[1:]
		} else {
			jumble = append(jumble, s[0])
			s = s[1:]
		}
	}

	return jumble
}

func main() {
	// read in from standard out
	line := "car truck 8 4 bus 6 1"
	r := JumbleSort(line)
	fmt.Printf("%v\n", r)
}
