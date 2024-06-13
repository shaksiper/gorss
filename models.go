package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/go-tutorial/internal/database"
)

type User struct {
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Name      string       `json:"name"`
	ApiKey    string       `json:"api_key"`
	ID        uuid.UUID    `json:"id"`
	Deleted   sql.NullBool `json:"deleted"`
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
