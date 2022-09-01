package app

type Config struct {
	MongoDbDsn        string
	ApiEndpoint       string
	FollowersEndpoint string
	FollowingEndpoint string
	FollowersPageSize string
	FollowingPageSize string
	InstagramCookies  string
	InstagramAppId    string
}

func InitConfig() Config {
	return Config{
		MongoDbDsn:        "",
		ApiEndpoint:       "https://i.instagram.com/api/v1/",
		FollowersEndpoint: "friendships/434045099/followers/",
		FollowingEndpoint: "friendships/434045099/following/",
		FollowersPageSize: "200",
		FollowingPageSize: "200",
		InstagramCookies:  "",
		InstagramAppId:    "",
	}
}
