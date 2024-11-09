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

func handleGetUserProfile(id int) (*UserProfile, error) {
	comments, err := getComments(id)
	if err != nil {
		return nil, err
	}
	likes, err := getLikes(id)
	if err != nil {
		return nil, err
	}
	friends, err := getFriends(id)
	if err != nil {
		return nil, err
	}
	return &UserProfile{
		ID:       id,
		Comments: comments,
		Likes:    likes,
		Friends:  friends,
	}, nil
}

func getComments(id int) ([]string, error) {
	time.Sleep(time.Millisecond * 200) // mimic http round trip
	comments := []string{
		"hey, that was great",
		"yeah buddy",
		"oh, I didn't know that",
	}
	return comments, nil
}

func getLikes(id int) (int, error) {
	time.Sleep(time.Millisecond * 200) // mimic http round trip
	return 33, nil
}

func getFriends(id int) ([]int, error) {
	time.Sleep(time.Millisecond * 100) // mimic http round trip
	return []int{11, 34, 854, 455}, nil
}
