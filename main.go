package main

import (
	"fmt"
	"log"
	"time"
)

// package context is a standard library package
func main() {
	start := time.Now()
	result, err := thirdpartyHTTPCall()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v:%+v\n", time.Since(start), result)
}

// mimic an http round trip
func thirdpartyHTTPCall() (string, error) {
	// mimic latency
	time.Sleep(time.Millisecond * 100)
	return "some response", nil
}
