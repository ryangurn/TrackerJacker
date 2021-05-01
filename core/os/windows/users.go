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
			return
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
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if uint32(us) == u.BadPasswordCount {
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