package models

import "time"

// User exported
type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

// Session Exported
type Session struct {
	UserName     string
	LastActivity time.Time
}
