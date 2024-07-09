package main

import (
	"fmt"
	"strings"
)

// 1. advanced interface mechanics
// 2. typed functions

// WHEN DO I USE TYPED FUNCTIONS AND WHEN DO I USE AN INTERFACE?
// - use an interface if you want to have state (a typed function has no way to store state)
type Foo interface{}

type FooImplementation struct {
	state any
}

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello Sailor!", Uppercase))
	fmt.Println(transformString("hello Sailor!", Prefixer("FOO_")))
	fmt.Println(transformString("hello Sailor!", Prefixer("BAR_")))
}
