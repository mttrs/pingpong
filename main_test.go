package main

import "testing"

func TestReserseStr(t *testing.T) {
	actual := ReverseStr("hello")
	expected := "olleh"

	if actual != expected {
		t.Fatalf("Expected %s but actual %s", expected, actual)
	}
}
