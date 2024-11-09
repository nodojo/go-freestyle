package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// package context is a standard library package
func main() {
	// we need a way to cancel third party calls when they don't return in a reasonable amount of time
	// -> this is where the use of context comes in
	start := time.Now()
	userID, err := fetchUserID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v:%+v\n", time.Since(start), userID)
}

func fetchUserID() (string, error) {
	// what is the maximum amount of time we are willing to wait for a response?
	// the cancel here allows us to manually cancel the call.
	// each time you create a context, you must provide a parent context...
	// if you don't have one, you can make a new context by using context.Background().
	// we are saying that any call that takes longer than 100ms will be canceled
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	// no matter what happens, if this function is going to return always cancel
	defer cancel()

	type result struct {
		userId string
		err    error
	}
	resultch := make(chan result, 1)

	// we are using a go routine, so we need to find a way to communicate the
	// values from this go routine to the main function
	// -> we do this with channels (see resultch defined above)
	go func() {
		res, err := thirdpartyHTTPCall()
		// write the result from this third party call into this channel
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	// there are two possibilities that can happen...
	// 1. we are going to have a response in the resultch, or
	// 2. the context is canceled
	select {
	// Done()
	// 1. context timeout is exceeded
	// 2. context has been manually canceled -> Cancel() (defer cancel() above)
	case <-ctx.Done():
		return "", ctx.Err() // return the error from the context
	case res := <-resultch:
		return res.userId, res.err
	}
}

// mimic an http round trip
// this represents any call where we can't control how long it's going to take
func thirdpartyHTTPCall() (string, error) {
	// mimic latency
	time.Sleep(time.Millisecond * 500)
	return "user id: 1", nil
}
