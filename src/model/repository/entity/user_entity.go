package entity

import "time"

type UserEntity struct {
	ID         int       `json:"id,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Name       string    `json:"name,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}
