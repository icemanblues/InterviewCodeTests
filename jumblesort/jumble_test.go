package main

import "testing"

func TestJumble(t *testing.T) {
	tableTests := []struct {
		name   string
		input  string
		output []string
	}{
		{
			name:   "Example 1",
			input:  "1",
			output: []string{"1"},
		},
		{
			name:   "Example 2",
			input:  "car truck bus",
			output: []string{"bus", "car", "truck"},
		},
		{
			name:   "Example 3",
			input:  "8 4 6 1 -2 9 5",
			output: []string{"-2", "1", "4", "5", "6", "8", "9"},
		},
		{
			name:   "Example 4",
			input:  "car truck 8 4 bus 6 1",
			output: []string{"bus", "car", "1", "4", "truck", "6", "8"},
		},
	}

	for _, test := range tableTests {
		actual := JumbleSort(test.input)
		if la, lo := len(actual), len(test.output); la != lo {
			t.Fatalf("%v: Not the correct length. actual: %v expected: %v", test.name, la, lo)
		}

		for i, a := range actual {
			if b := test.output[i]; a != b {
				t.Errorf("%v: The element at index %v does not match. actual: %v expected: %v", test.name, i, a, b)
			}
		}
	}
}
