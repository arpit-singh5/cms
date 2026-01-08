package models

type User struct {
	ID        uint
	Email     string
	Name      string
	HashPass  string
	Bio       string
	Image_Uri string
	SocialUri []string
}

type Role struct {
	ID   uint
	Name string
}

type Permission struct {
	ID   uint
	Name string
}
