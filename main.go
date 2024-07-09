package main

import (
	"fmt"
	"strings"
)

// 1. advanced interface mechanics
// 2. typed functions

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// this won't work, because TransformFunc only accepts a single string param
func Prefixer(s string, prefix string) string {
	return prefix + s
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello Sailor!", Uppercase))
	fmt.Println(transformString("hello Sailor!", Prefixer))
}
