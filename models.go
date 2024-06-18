package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/go-tutorial/internal/database"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	ID        uuid.UUID `json:"id"`
	Deleted   bool      `json:"deleted"`
}

// Map database user to User DTO
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Deleted:   dbUser.Deleted,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Deleted   bool      `json:"deleted"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		ID:        dbFeed.ID,
		UserID:    dbFeed.UserID,
		Deleted:   dbFeed.Deleted,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		ID:        dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
		FeedId:    dbFeedFollow.FeedID,
	}
}
