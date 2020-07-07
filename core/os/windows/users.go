package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"os/user"
)

func WAPIUserExist(usr string) (retBool bool) {
	retBool = false

	users, err := wapi.ListLocalUsers()
	if err != nil {
		fmt.Printf("Error fetching user list, %s.\r\n", err.Error())
		return
	}

	for _, u := range users {
		if usr == u.Username {
			retBool = true
		}
	}

	return
}

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
		if WAPIUserExist(args[1]) == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}