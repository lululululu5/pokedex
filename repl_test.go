package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	actual := cleanInput("Hello World")
	fmt.Println(actual)
	expected := []string{
		"hello",
		"world",
	}

	if len(actual) != len(expected) {
		t.Errorf("lengths don't match: %v vs %v", actual, expected)
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Words don't match")
		}

	}

} 