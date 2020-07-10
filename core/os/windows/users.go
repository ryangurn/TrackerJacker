package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"strings"
	"time"
)

func UserExist(usr string) (retBool bool) {
	retBool = false

	users, err := wapi.ListLocalUsers()
	if err != nil {
		fmt.Printf("Error fetching user list, %s.\r\n", err.Error())
		return
	}

	for _, u := range users {
		if strings.TrimSpace(strings.ToLower(usr)) == strings.TrimSpace(strings.ToLower(u.Username)) {
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
				if value.(uint32) == u.BadPasswordCount {
					retBool = true
					return
				}
			} else if key == "FullName" {
				if value.(string) == u.FullName {
					retBool = true
					return
				}
			} else if key == "IsAdmin" {
				if value.(string) == strconv.FormatBool(u.IsAdmin) {
					retBool = true
					return
				}
			} else if key == "IsEnabled" {
				if value.(string) == strconv.FormatBool(u.IsEnabled) {
					retBool = true
					return
				}
			} else if key == "IsLocked" {
				if value.(string) == strconv.FormatBool(u.IsLocked) {
					retBool = true
					return
				}
			} else if key == "LastLogon" {
				if value.(time.Time) == u.LastLogon {
					retBool = true
					return
				}
			} else if key == "NoChangePassword" {
				if value.(string) == strconv.FormatBool(u.NoChangePassword) {
					retBool = true
					return
				}
			} else if key == "NumberOfLogons" {
				if value.(uint32) == u.NumberOfLogons {
					retBool = true
					return
				}
			} else if key == "PasswordAge" {
				if value.(time.Duration) == u.PasswordAge {
					retBool = true
					return
				}
			} else if key == "PasswordNeverExpires" {
				if value.(string) == strconv.FormatBool(u.PasswordNeverExpires) {
					retBool = true
					return
				}
			}
		}
	}

	return
}

func UserLoggedIn(usr string) (retBool bool) {
	retBool = false

	users, err := wapi.ListLoggedInUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if strings.TrimSpace(strings.ToLower(usr)) == strings.TrimSpace(strings.ToLower(u.Username)) {
			retBool = true
		}
	}

	return
}

func UserParse(args []string, result interface{}) (retBool bool) {
	retBool = false


	if args[0] == "exist" {
		if len(args) != 2 {
			return
		}
		if UserExist(args[1]) == result {
			retBool = true
		}
	} else if args[0] == "meta" {
		if len(args) != 4 {
			return
		}
		if UserMeta(args[1], args[2], args[3]) == result {
			retBool = true
		}
	} else if args[0] == "logged" {
		if len(args) != 3 {
			return
		}

		if args[1] == "exist" {
			if UserLoggedIn(args[2]) == result {
				retBool = true
			}
		} else {
			fmt.Printf("Unrecognized Command: %s\n", args[0])
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}