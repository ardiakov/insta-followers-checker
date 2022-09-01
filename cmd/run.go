package main

import (
	"fmt"
	"insta-follower-notifier/internal/app"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	users := context.InstagramClient.GetFollowers()

	for _, user := range users {
		println(user)
	}
}
