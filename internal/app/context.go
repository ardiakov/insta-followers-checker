package app

import (
	"insta-follower-notifier/internal/instagram"
	"insta-follower-notifier/internal/mongodb"
)

type Context struct {
	InstagramClient *instagram.InstagramClient
	MongoDbClient   *mongodb.MongoDbClient
}

func InitContext(config Config) *Context {
	context := &Context{
		InstagramClient: instagram.InitClient(
			config.ApiEndpoint,
			config.FollowersEndpoint,
			config.FollowingEndpoint,
			config.FollowersPageSize,
			config.FollowingPageSize,
			config.InstagramCookies,
			config.InstagramAppId,
		),
		MongoDbClient: mongodb.InitClient(),
	}

	return context
}
