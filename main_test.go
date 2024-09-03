package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}
	have := Player{
		name: "Alice",
		hp:   59,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}
}

func TestCalculateValuesSpecial(t *testing.T) {
	var (
		expected = 10
		a        = 4
		b        = 5
	)
	// have is the value I get back from the test
	have := calculateValues(a, b)
	if have != expected {
		// Errorf() -> print a formatted error
		// using %d because we are working with integers
		t.Errorf("expected %d but have %d", expected, have)
	}
}
