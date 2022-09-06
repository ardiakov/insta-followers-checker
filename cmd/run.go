package main

import (
	"fmt"
	"insta-follower-notifier/internal/app"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	//followers := context.InstagramClient.GetFollowers()
	//context.MongoDbClient.Update("followers", "followers", followers)
	//
	//followings := context.InstagramClient.GetFollowings()
	//context.MongoDbClient.Update("followers", "followings", followings)

	diff := context.MongoDbClient.GetDiff("followers", "followings")

	fmt.Println(diff)

}
