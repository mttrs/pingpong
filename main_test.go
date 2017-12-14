package main

import "testing"

var tests = []struct {
	in  string
	out string
}{
	{"hello", "olleh"},
	{"Hi", "iH"},
	{"おはよう", "うよはお"},
}

func TestReserseStr(t *testing.T) {
	for _, tt := range tests {
		actual := ReverseStr(tt.in)
		expected := tt.out

		if actual != expected {
			t.Fatalf("Expected %s but actual %s", expected, actual)
		}
	}
}
