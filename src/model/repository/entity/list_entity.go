package entity

import "time"

type ListEntity struct {
	ID          int
	Title       string
	Description string
	UserID      int
	Created_at  time.Time
	Deleted_at  time.Time
	Updated_at  time.Time
}
