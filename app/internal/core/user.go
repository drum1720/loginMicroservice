package core

import (
	"errors"
	"time"
)

type User struct {
	UserId    int64  `json:"user_id"`
	User      string `json:"user"`
	Password  string `json:"password"`
	LastVisit time.Time
}

func (u User) Validation() error {
	if u.User == "" {
		return errors.New("user needed")
	}
	if len(u.Password) < 4 {
		return errors.New("short password")
	}
	return nil
}

func NewUser() *User {
	return &User{}
}
