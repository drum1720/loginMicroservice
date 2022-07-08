package core

import (
	"errors"
	"time"
)

type User struct {
	UserId    int64
	User      string `json:"user"`
	Password  string `json:"password"`
	LastVisit time.Time
}

// Validate ...
func (u User) Validate() error {
	if u.User == "" {
		return errors.New("user needed")
	}
	if len(u.Password) < 4 {
		return errors.New("short password")
	}
	return nil
}
