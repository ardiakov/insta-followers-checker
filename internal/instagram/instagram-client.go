package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Instagram struct {
	Endpoint          string
	followersEndpoint string
	followingEndpoint string
	FollowersPageSize string
	FollowingPageSize string
}

func InitClient() *Instagram {
	return &Instagram{
		Endpoint:          "https://i.instagram.com/api/v1/",
		followersEndpoint: "friendships/434045099/followers/",
		followingEndpoint: "friendships/434045099/following/",
		FollowersPageSize: "200",
	}
}

type FollowersResponse struct {
	Users []struct {
		Pk       int    `json:"pk"`
		Username string `json:"username"`
		FullName string `json:"full_name"`
	} `json:"users"`
	PageSize  int    `json:"page_size"`
	NextMaxId string `json:"next_max_id"`
	Status    string `json:"status"`
}

type FollowingResponse struct {
	Users []struct {
		Pk       int    `json:"pk"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	} `json:"users"`
	PageSize  int    `json:"page_size"`
	NextMaxId string `json:"next_max_id"`
	Status    string `json:"status"`
}

func (r *Instagram) GetFollowers() {
	counter := 0
	maxId := ""

	for {
		parm := url.Values{}
		parm.Add("count", r.FollowersPageSize)

		if maxId != "" {
			parm.Add("max_id", maxId)
		}

		req, err := http.NewRequest("GET", r.Endpoint+r.followersEndpoint+"?", nil)
		req.Header.Set("cookie", os.Getenv("COOKIE"))
		req.Header.Set("x-ig-app-id", os.Getenv("APP_ID"))

		req.URL.RawQuery = parm.Encode()

		if err != nil {
			panic(err)
		}

		resp, err := http.DefaultClient.Do(req)

		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		data := FollowersResponse{}

		_ = json.Unmarshal(bytes, &data)

		for _, user := range data.Users {
			counter++
			fmt.Println(strconv.Itoa(counter) + " " + user.Username)
		}

		maxId = data.NextMaxId

		if maxId == "" {
			return
		}
	}
}
func (r *Instagram) GetFollowing() {
	maxId := ""

	parm := url.Values{}
	parm.Add("count", string(r.FollowingPageSize))
	parm.Add("max_id", maxId)

	fmt.Println(parm.Encode())

	req, err := http.NewRequest("GET", r.Endpoint+r.followingEndpoint+"?", strings.NewReader(parm.Encode()))
	req.Header.Set("cookie", os.Getenv("COOKIE"))
	req.Header.Set("x-ig-app-id", os.Getenv("APP_ID"))

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Status)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := FollowingResponse{}

	_ = json.Unmarshal(bytes, &data)

	nexMaxId := data.NextMaxId
	//if nexMaxId  {
	//	panic("Undefinex next_max_id")
	//}

	for _, user := range data.Users {
		fmt.Println(string(user.Pk) + " " + user.Username)
	}

	fmt.Println(nexMaxId)
}
