package main

import (
	"fmt"
	"insta-follower-notifier/internal/app"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	//context.MongoDbClient.CreateDocument()

	//followers := context.InstagramClient.GetFollowers()
	//context.MongoDbClient.Update("followers", "followers", followers)
	//followings := context.InstagramClient.GetFollowings()
	//context.MongoDbClient.Update("followers", "followings", followings)

	context.MongoDbClient.GetDiff("followers", "followings")

}
