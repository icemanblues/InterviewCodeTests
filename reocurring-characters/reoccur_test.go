package main

import "testing"

func TestReoccur(t *testing.T) {
	tests := []struct {
		input string
		r     rune
		ok    bool
	}{
		{"abcdefghiabc", 'a', true},
		{"cc", 'c', true},
		{"", 'd', false},
		{"xylem", 'x', false},
		{"cabbage", 'b', true},
		{"12345 098765 blah", '5', true},
	}

	for _, test := range tests {
		t.Run(test.input, func(tt *testing.T) {
			actual, ok := Reoccur(test.input)
			if ok != test.ok {
				tt.Errorf("Incorrect ok. Expected: %v Actual: %v", test.ok, ok)
			}
			if ok && test.r != actual {
				tt.Errorf("Incorrect Rune. Expected: %c Actual: %c", test.r, actual)
			}
		})
	}
}

func TestReoccurNoMem(t *testing.T) {
	tests := []struct {
		input string
		r     rune
		ok    bool
	}{
		{"abcdefghiabc", 'a', true},
		{"cc", 'c', true},
		{"", 'd', false},
		{"xylem", 'x', false},
		{"cabbage", 'b', true},
		{"12345 098765 blah", '5', true},
	}

	for i, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual, ok := ReoccurNoMem(test.input)
			if ok != test.ok {
				t.Errorf("%v: %v | Incorrect ok. Expected: %v Actual: %v", i, test.input, test.ok, ok)
			}
			if ok && actual != test.r {
				t.Errorf("%v: %v | Incorrect rune. Expected: %c Actual: %c", i, test.input, test.r, actual)
			}
		})
	}
}
