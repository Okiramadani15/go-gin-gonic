package models

import "time"

type User struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Address string `json:"address"`

	Birthdate time.Time `json:"birthdata"`
}
