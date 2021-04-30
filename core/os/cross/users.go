package cross

import (
	"os/user"
)

// TODO: Add support for meta and logged in checks

func UserExist(usr string) (retBool bool) {
	retBool = false
	u, err := user.Lookup(usr)

	if err != nil {
		return
	}

	if u != nil {
		retBool = true
	}

	return
}