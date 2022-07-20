package security

import (
	"fmt"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	token, err := CreateJWT("site", "ss1941-1945HG_bl")
	if err != nil {

	}
	dep, err := ParseJWT(token, "11546r75675675675")
	fmt.Println(dep)
}
