package client_repo

import (
	"time"
)

type Model struct {
	ID          int    `json:"-" db:"id"`
	UUID        string `json:"uuid" db:"uuid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	// Status flags
	Deleted  bool `json:"deleted" db:"deleted"`
	Archived bool `json:"archived" db:"archived"`

	// Address fields
	AddressLine1 string `json:"address_line1" db:"address_line1"`
	AddressLine2 string `json:"address_line2,omitempty" db:"address_line2"`
	City         string `json:"city" db:"city"`
	PostalCode   string `json:"postal_code" db:"postal_code"`
	Country      string `json:"country" db:"country"`

	// Geolocation fields
	Latitude  *float64 `json:"latitude,omitempty" db:"latitude"`
	Longitude *float64 `json:"longitude,omitempty" db:"longitude"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ModelWithDistance struct {
	Model
	Distance float64 `json:"distance_miles" db:"distance"`
}
