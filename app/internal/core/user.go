package core

import (
	"errors"
	"time"
)

type User struct {
	UserId    int64     `json:"-"`
	User      string    `json:"user"`
	Password  string    `json:"password"`
	LastVisit time.Time `json:"last_visit"`
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
