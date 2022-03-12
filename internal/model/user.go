package model

import "time"

type UserRegisterRequest struct {
	Name      string
	Birthdate time.Time
}

type UserRegisterResponse struct {
	Name      string    `json:"full_name"`
	Birthdate time.Time `json:"birth_date"`
}
