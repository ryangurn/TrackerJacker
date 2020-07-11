package cross

import "os/user"

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

func UserParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 2 {
		return
	}

	if args[0] == "exist" {
		if UserExist(args[1]) == result {
			retBool = true
			return
		}
	}

	return
}