package model

import "time"

type Book struct {
	Id        int
	Title     string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
