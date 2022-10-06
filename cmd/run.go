package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"insta-follower-notifier/internal/app"
	"log"
	"time"
)

func main() {
	fmt.Println("Run")

	config := app.InitConfig()
	context := app.InitContext(config)

	context.MongoDbClient.CreateDocument()

	unfollowers := context.MongoDbClient.GetUnfollowers()

	followers := context.InstagramClient.GetFollowers()

	time.Sleep(time.Second * 5)
	context.MongoDbClient.Update("followers", "followers", followers)

	followings := context.InstagramClient.GetFollowings()

	time.Sleep(time.Second * 5)
	context.MongoDbClient.Update("followers", "followings", followings)

	context.MongoDbClient.UpdateUnfollowers("followers", "followings")
	newUnfollowers := context.MongoDbClient.GetUnfollowers()

	currentSheet := time.Now().Format("2006-01-02")

	file := excelize.NewFile()
	file.NewSheet(currentSheet)
	file.DeleteSheet("Sheet1")

	err := file.SetCellValue(currentSheet, "A1", fmt.Sprintf("Отписавшиеся: %d", len(newUnfollowers)))

	if err != nil {
		log.Fatal(err)
	}

	for i, unfollower := range newUnfollowers {
		err = file.SetCellValue(currentSheet, fmt.Sprintf("A%d", i+2), unfollower)

		if err != nil {
			log.Fatal(err)
		}
	}

	diffUsers := context.MongoDbClient.DiffBetweenUnfollowers(unfollowers, newUnfollowers)

	err = file.SetCellValue(currentSheet, "B1", fmt.Sprintf("Новые отписавшиеся: %d", len(diffUsers)))

	if err != nil {
		log.Fatal(err)
	}

	for i, newDiffUser := range diffUsers {
		err = file.SetCellValue(currentSheet, fmt.Sprintf("B%d", i+2), newDiffUser)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err = file.SaveAs("result.xlsx"); err != nil {
		log.Fatal(err)
	}
}
