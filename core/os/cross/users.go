package cross

import "os/user"

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
