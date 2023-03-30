package app

type Config struct {
	MongoDbDsn        string
	MongoDbName       string
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
		MongoDbDsn:        "mongodb://mongo:27017",
		MongoDbName:       "insta",
		ApiEndpoint:       "https://i.instagram.com/api/v1/",
		FollowersEndpoint: "friendships/434045099/followers/",
		FollowingEndpoint: "friendships/434045099/following/",
		FollowersPageSize: "200",
		FollowingPageSize: "200",
		InstagramCookies:  "csrftoken=y8l7LUCNXcP2jxTVXNcFVwX0NePsCNhi; ds_user_id=434045099; rur=\"LDC\\054434045099\\0541695223800:01f7e59cc6a9cafe2d13418187a4dd442b79a206afec6760e0b6554b669dfe0b49ea3ebe\"; dpr=2; datr=stspYzDPPiL6jc2FbpptX7f3; shbid=\"8142\\054434045099\\0541695223602:01f70db54d8ee792df0d20eb19839480ae55a80349e6158505e48fac57e8456e21aadb04\"; shbts=\"1663687602\\054434045099\\0541695223602:01f71a91f841994cfef0e3a308f1276a38bf3c25ed554ebed334a1fe1fc34bf8228612b2\"; sessionid=434045099%3AY0QHb0KqBGtR4B%3A10%3AAYe1jmHHTgCf_Xbes2M4H8LbgMWBYxAMSvRWyHqbWA; ig_did=D50EA7C6-27B5-439D-B334-92E6DDFFA794; mid=YynbnwAEAAGS4r6Rci_i7w0Vxgml",
		InstagramAppId:    "936619743392459",
	}
}
