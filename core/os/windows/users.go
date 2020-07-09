package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
)

func UserExist(usr string) (retBool bool) {
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

func UserMeta(usr string, key string, value interface{}) (retBool bool) {
	retBool = false

	users, err := wapi.ListLocalUsers()
	if err != nil {
		fmt.Printf("Error fetching user list, %s\r\n", err.Error())
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if key == "BadPasswordCount" {
				if value == u.BadPasswordCount {
					retBool = true
					return
				}
			} else if key == "FullName" {
				if value == u.FullName {
					retBool = true
					return
				}
			} else if key == "IsAdmin" {
				if value == u.IsAdmin {
					retBool = true
					return
				}
			} else if key == "IsEnabled" {
				if value == u.IsEnabled {
					retBool = true
					return
				}
			} else if key == "IsLocked" {
				if value == u.IsLocked {
					retBool = true
					return
				}
			} else if key == "LastLogin" {
				if value == u.LastLogon {
					retBool = true
					return
				}
			} else if key == "NoChangePassword" {
				if value == u.NoChangePassword {
					retBool = true
					return
				}
			} else if key == "NumberOfLogons" {
				if value == u.NumberOfLogons {
					retBool = true
					return
				}
			} else if key == "PasswordAge" {
				if value == u.PasswordAge {
					retBool = true
					return
				}
			} else if key == "PasswordNeverExpires" {
				if value == u.PasswordNeverExpires {
					retBool = true
					return
				}
			}
		}
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
	} else if args[0] == "meta" {
		if UserMeta(args[1], "IsAdmin", true) == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}