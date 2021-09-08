package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"time"
)

func UserExist(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			retBool = true
			if out, err := json.Marshal(u); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserLoggedIn(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLoggedInUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			retBool = true
			if out, err := json.Marshal(u); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserBadPassword(usr string, count string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	val, err := strconv.ParseUint(count, 10, 64)
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.BadPasswordCount == uint32(val) {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserFullName(usr string, name string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.FullName == name {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserAdmin(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.IsAdmin == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserEnabled(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.IsEnabled == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserLocked(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.IsLocked == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserLastLogon(usr string, date string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	val, err := time.Parse("2006-01-02 15:04", date)
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.LastLogon == val {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserNoChangePassword(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.NoChangePassword == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserPasswordChangeable(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.NoChangePassword == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserNoOfLogons(usr string, count string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	val, err := strconv.ParseUint(count, 10, 64)
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.NumberOfLogons == uint32(val) {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserPasswordAge(usr string, duration string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	val, err := time.ParseDuration(duration)
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.PasswordAge == val {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}

func UserPasswordExpires(usr string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	users, err := wapi.ListLocalUsers()
	if err != nil {
		return
	}

	for _, u := range users {
		if usr == u.Username {
			if u.PasswordNeverExpires == true {
				retBool = true
				if out, err := json.Marshal(u); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(users)
	return retBool, string(out)
}