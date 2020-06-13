package domain

import "time"

type Profile struct {
	ID           int64
	Name         string
	Introduction string
	Sex          Sex
	UserID       int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Sex int64

const (
	Male Sex = iota + 1
	Female
)
