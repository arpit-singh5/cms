package models

import "time"

type PostStatus string

const (
	StatusDraft     PostStatus = "draft"
	StatusPublished PostStatus = "published"
	StatusArchived  PostStatus = "archived"
)

type Post struct {
	ID            uint
	UserId        uint
	Title         string
	Slug          string
	Description   string
	Content       string
	FeaturedImage string
	Status        PostStatus
	CreatedAt     time.Time
	PublishedAt   time.Time
	UpdatedAt     time.Time
}

func (s PostStatus) IsValid() bool {
	switch s {
	case StatusArchived, StatusDraft, StatusPublished:
		return true
	default:
		return false

	}
}

type Tag struct {
	ID   uint
	Name string
	Slug string
}
