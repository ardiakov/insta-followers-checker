package main

import (
	"fmt"
	"insta-follower-notifier/internal/app"
	"time"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	//context.MongoDbClient.CreateDocument()

	unfollowers := context.MongoDbClient.GetUnfollowers()

	followers := context.InstagramClient.GetFollowers()

	time.Sleep(time.Second * 5)
	context.MongoDbClient.Update("followers", "followers", followers)

	followings := context.InstagramClient.GetFollowings()
	time.Sleep(time.Second * 5)
	context.MongoDbClient.Update("followers", "followings", followings)

	context.MongoDbClient.UpdateUnfollowers("followers", "followings")
	newUnfollowers := context.MongoDbClient.GetUnfollowers()

	fmt.Println(newUnfollowers)

	users := context.MongoDbClient.DiffBetweenUnfollowers(unfollowers, newUnfollowers)

	fmt.Println(users)
}
