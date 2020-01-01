package domain

import "time"

type Profile struct {
	ID        int64
	Name      string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
