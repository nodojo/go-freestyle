package main

import (
	"fmt"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	// this won't print when we run "go test ./..."
	fmt.Println("hello from our test")
	// to print the above line when testing, run "go test -v ./..."
}
