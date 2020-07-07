package windows

import (
	"fmt"
	"os/user"
)

func UserExist(usr string) (retBool bool) {
	retBool = false
	user, err := user.Lookup(usr)

	if err != nil {
		return
	}

	if user != nil {
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
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}