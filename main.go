package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userProfile)
	fmt.Println("fetching user profile took ", time.Since(start))
}

// we can make this better by using goroutines to fetch
// the different parts of a user profile asynchronously
type Response struct {
	data any
	err  error
}

func handleGetUserProfile() (*UserProfile, error) {
	// create channel and buffer for the 3 responses (comments, likes, friends).
	// coordinate workers (goroutines) so we will know when they are done
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	// schedule this in goroutines.
	// we are making 3 requests, each inside its own goroutine...
	go getComments(respch, wg)
	go getLikes(respch, wg)
	go getFriends(respch, wg)
	// ...this means, we need to add 3 to the waitgroup
	wg.Add(3)
	// block until the wg counter == 0, then unblock
	wg.Wait()
	// now that wg is unblocked, close the channel
	close(respch) // this signals the range below to break out of the loop

	// range over our response channel
	userProfile := &UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}

		// switch on the type of data being sent
		// note: data types can be replaced with custom data types for cleaner logic
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg // automatically casts to int
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

// we won't need the return because we are going to inject the data into the channel
func getComments(respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200) // mimic http round trip
	comments := []string{
		"hey, that was great",
		"yeah buddy",
		"oh, I didn't know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}

	// each time worker is done, signal completion by calling wg.Done()
	wg.Done()
}

// we won't need the return because we are going to inject the data into the channel
func getLikes(respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200) // mimic http round trip
	respch <- Response{
		data: 11,
		err:  nil,
	}

	// each time worker is done, signal completion by calling wg.Done()
	wg.Done()
}

// we won't need the return because we are going to inject the data into the channel
func getFriends(respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100) // mimic http round trip
	friendIds := []int{11, 34, 854, 455}
	respch <- Response{
		data: friendIds,
		err:  nil,
	}

	// each time worker is done, signal completion by calling wg.Done()
	wg.Done()
}
