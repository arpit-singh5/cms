package models

import "time"

type Post struct {
	ID uint
	UserId uint
	Title string
	Slug string
	Description string
	Content string
	FeaturedImage string
	CreatedAt time.Time
}

type Tag struct {
	ID uint
	Name string
	Slug string
}
