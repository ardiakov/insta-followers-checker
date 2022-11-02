package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type InstagramClient struct {
	Endpoint          string
	followersEndpoint string
	followingEndpoint string
	followersPageSize string
	followingPageSize string
	instagramCookies  string
	instagramAppId    string
}

func InitClient(
	apiEndpoint string,
	followersEndpoint string,
	followingEndpoint string,
	followersPageSize string,
	followingPageSize string,
	instagramCookies string,
	instagramAppId string) *InstagramClient {
	return &InstagramClient{
		Endpoint:          apiEndpoint,
		followersEndpoint: followersEndpoint,
		followingEndpoint: followingEndpoint,
		followersPageSize: followersPageSize,
		followingPageSize: followingPageSize,
		instagramCookies:  instagramCookies,
		instagramAppId:    instagramAppId,
	}
}

type FollowersResponse struct {
	Users []struct {
		Pk       string `json:"pk"`
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

func (r *InstagramClient) GetFollowers() []string {
	fmt.Println("Получаем список подписчиков")

	maxId := ""

	var users []string

	for {
		parm := url.Values{}
		parm.Add("count", r.followersPageSize)
		parm.Add("search_surface", "follow_list_page")

		if maxId != "" {
			parm.Add("max_id", maxId)
		}

		req, err := http.NewRequest("GET", r.Endpoint+r.followersEndpoint+"?", nil)

		req.Header.Set("cookie", r.instagramCookies)
		req.Header.Set("x-ig-app-id", r.instagramAppId)

		req.URL.RawQuery = parm.Encode()

		if err != nil {
			panic(err)
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		data := FollowersResponse{}

		err = json.Unmarshal(bytes, &data)

		if err != nil {
			panic(err)
		}

		for _, user := range data.Users {
			//counter++
			//fmt.Println(strconv.Itoa(counter) + " " + user.Username)
			users = append(users, user.Username)
		}

		maxId = data.NextMaxId

		if maxId == "" {
			break
		}
	}

	return users
}
func (r *InstagramClient) GetFollowings() []string {
	fmt.Println("Получаем список подписок")

	counter := 0
	maxId := ""

	var users []string

	for {
		parm := url.Values{}
		parm.Add("count", r.followingPageSize)

		if maxId != "" {
			parm.Add("max_id", maxId)
		}

		req, err := http.NewRequest("GET", r.Endpoint+r.followingEndpoint+"?", nil)
		req.Header.Set("cookie", r.instagramCookies)
		req.Header.Set("x-ig-app-id", r.instagramAppId)

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

		data := FollowingResponse{}

		_ = json.Unmarshal(bytes, &data)

		for _, user := range data.Users {
			counter++
			users = append(users, user.Username)
		}

		maxId = data.NextMaxId

		if maxId == "" {
			break
		}
	}

	return users
}
