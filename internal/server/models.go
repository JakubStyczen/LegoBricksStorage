package server

import (
	"time"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

type LegoSet struct {
	ID           int32  `json:"id"`
	SerialNumber string `json:"serial_number"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	Theme        string `json:"theme"`
	Year         int32  `json:"year"`
	TotalParts   int32  `json:"total_parts"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Age       int32     `json:"age"`
	ApiKey    string    `json:"api_key"`
}

func databaseLegoSetToLegoSet(lego_set database.LegoSet) LegoSet {
	return LegoSet{
		ID:           lego_set.ID,
		SerialNumber: lego_set.SerialNumber,
		Name:         lego_set.Name,
		Price:        lego_set.Price,
		Theme:        lego_set.Theme,
		Year:         lego_set.Year,
		TotalParts:   lego_set.TotalParts,
	}
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Age:       user.Age,
		ApiKey:    user.ApiKey,
	}
}
