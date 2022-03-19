package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Instagram struct {
	Endpoint          string
	followersEndpoint string
	FollowersPageSize int
}

func InitClient() *Instagram {
	return &Instagram{
		Endpoint:          "https://i.instagram.com/api/v1/",
		followersEndpoint: "friendships/434045099/followers/",
		FollowersPageSize: 20,
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

func (r *Instagram) GetFollowers() {
	maxId := "QVFCS1VvYURINnlwQm5WQnBrTm1TT3Z5WWZqY0tnNzlLVTQzY2JERkhNVThqaU1oa2lKcU1aQ3AxOXpBb3ZxbER2aHJZWmhwZFI5d2dTV0x1R3ZEdDNIMQ%3D%3D"

	parm := url.Values{}
	parm.Add("count", string(r.FollowersPageSize))
	parm.Add("max_id", maxId)

	fmt.Println(parm.Encode())

	req, err := http.NewRequest("GET", r.Endpoint+r.followersEndpoint+"?", strings.NewReader(parm.Encode()))
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

	data := FollowersResponse{}

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
