package request

import "errors"

type SecretSequrDangerWork struct {
	User            string `json:"user"`
	AutorizeWithJWT bool   `json:"-"`
}

func (r *SecretSequrDangerWork) Validate() error {
	if r.User == "" {
		return errors.New("user needed")
	}

	return nil
}
