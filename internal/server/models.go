package server

import (
	"time"

	"github.com/JakubStyczen/LegoBricksStorage/internal/database"
	"github.com/google/uuid"
)

type LegoSetPart struct {
	SetSerial  string `json:"set_serial_number"`
	PartSerial string `json:"part_serial_number"`
	Quantity   int32  `json:"quantity"`
}

type LegoPart struct {
	ID           uuid.UUID `json:"id"`
	SerialNumber string    `json:"serial_number"`
	Quantity     int32     `json:"quantity"`
	Name         string    `json:"name"`
	Color        string    `json:"color"`
	Price        string    `json:"price"`
}

type UserSet struct {
	ID      uuid.UUID `json:"id"`
	OwnedAt time.Time `json:"owned_at"`
	Price   string    `json:"price"`
	SetID   uuid.UUID `json:"set_id"`
	UserID  uuid.UUID `json:"user_id"`
}

type LegoSet struct {
	ID           uuid.UUID `json:"id"`
	SerialNumber string    `json:"serial_number"`
	Name         string    `json:"name"`
	Price        string    `json:"price"`
	Theme        string    `json:"theme"`
	Year         int32     `json:"year"`
	TotalParts   int32     `json:"total_parts"`
	UserID       uuid.UUID `json:"user_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Age       int32     `json:"age"`
	ApiKey    string    `json:"api_key"`
}

type LegoSetPartWithName struct {
	SetSerial  string `json:"set_serial_number"`
	PartSerial string `json:"part_serial_number"`
	SetName    string `json:"set_name"`
	PartName   string `json:"part_name"`
	Quantity   int32  `json:"quantity"`
}

func databaseLegoAllPartsInSetsToLegoSetPartWithName(lego_set_part database.GetAllPartsInAllSetsRow) LegoSetPartWithName {
	return LegoSetPartWithName{
		SetSerial:  lego_set_part.SetSerial,
		PartSerial: lego_set_part.PartSerial,
		SetName:    lego_set_part.SetName,
		PartName:   lego_set_part.PartName,
		Quantity:   lego_set_part.Quantity,
	}
}

func databaseLegoSetPartToLegoSetPart(lego_set_part database.LegoSetPart) LegoSetPart {
	return LegoSetPart{
		SetSerial:  lego_set_part.SetSerial,
		PartSerial: lego_set_part.PartSerial,
		Quantity:   lego_set_part.Quantity,
	}
}

func databaseLegoPartToLegoPart(lego_part database.LegoPart) LegoPart {
	return LegoPart{
		ID:           lego_part.ID,
		SerialNumber: lego_part.SerialNumber,
		Quantity:     lego_part.Quantity,
		Name:         lego_part.Name,
		Color:        lego_part.Color,
		Price:        lego_part.Price,
	}
}

func databaseUserSetToUserSet(user_set database.UserSet) UserSet {
	return UserSet{
		ID:      user_set.ID,
		OwnedAt: user_set.OwnedAt,
		Price:   user_set.Price,
		SetID:   user_set.SetID,
		UserID:  user_set.UserID,
	}
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
		UserID:       lego_set.UserID,
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
