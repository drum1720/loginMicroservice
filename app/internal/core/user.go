package core

import "time"

type User struct {
	UserId    int64
	Login     string
	Pass      string
	LastVisit time.Time
}

func NewUser() *User {
	return &User{}
}
