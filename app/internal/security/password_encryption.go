package security

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = "shla_kto-to_kudato_i_sossala~0~"
	cost = 6
)

func PassEncryption(pass string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(salt+pass), cost)
	if err != nil {
		return nil
	}

	return hashedPassword
}

func PassCorrect(pass string, hashedPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(salt+pass)); err != nil {
		return false
	}

	return true
}
