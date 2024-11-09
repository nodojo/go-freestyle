package main

import (
	"fmt"
	"log"
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
	userProfile, err := handleGetUserProfile(10)
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

func handleGetUserProfile(id int) (*UserProfile, error) {
	// create channel and buffer for the 3 responses (comments, likes, friends)
	respch := make(chan Response, 3)

	// schedule this in a goroutine
	go getComments(id, respch)
	go getLikes(id, respch)
	go getFriends(id, respch)

	// range over our response channel
	// DEADLOCK -> because range "keeps ranging" and doesn't know when to exit
	for resp := range respch {
		fmt.Println(resp)
	}

	return &UserProfile{
		// ID:       id,
		// Comments: comments,
		// Likes:    likes,
		// Friends:  friends,
	}, nil
}

// we won't need the return because we are going to inject the data into the channel
func getComments(id int, respch chan Response) {
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
}

// we won't need the return because we are going to inject the data into the channel
func getLikes(id int, respch chan Response) {
	time.Sleep(time.Millisecond * 200) // mimic http round trip
	respch <- Response{
		data: 11,
		err:  nil,
	}
}

// we won't need the return because we are going to inject the data into the channel
func getFriends(id int, respch chan Response) {
	time.Sleep(time.Millisecond * 100) // mimic http round trip
	friendIds := []int{11, 34, 854, 455}
	respch <- Response{
		data: friendIds,
		err:  nil,
	}
}
