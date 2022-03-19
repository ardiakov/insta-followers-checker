package main

import (
	"fmt"
	"insta-follower-notifier/internal/instagram"
)

func main() {
	fmt.Println("test")

	ints := instagram.InitClient()
	ints.GetFollowers()
}
