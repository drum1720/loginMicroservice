package request

import "errors"

type Authorize struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (r *Authorize) Validate() error {
	if r.User == "" {
		return errors.New("user needed")
	}
	if len(r.Password) < 4 {
		return errors.New("short password")
	}
	return nil
}
