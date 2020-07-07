package windows

import (
	"strconv"
	"strings"
	"time"

	wapi "github.com/iamacarpet/go-win64api"
	so "github.com/iamacarpet/go-win64api/shared"
)
type User []so.LocalUser

func SearchUser(so User, searchName string) (sw so.LocalUser) {
	for i := 0; i < len(so); i++ {
		if so[i].Username != "" && strings.Contains(strings.ToLower(so[i].Username), strings.ToLower(searchName)) {
			return so[i]
		}
	}

	return sw
}

func ExistUser(so User, searchName string) (retBool bool) {
	retBool = false
	for i := 0; i < len(so); i++ {
		if !retBool {
			if so[i].Username != "" && strings.Contains(strings.ToLower(so[i].Username), strings.ToLower(searchName)) {
				retBool = true
			}
		}
	}

	return
}

//UserCheckValue (This function will tell you first if a record exists in the users list and then allows you to check for a specific value search)
func UserCheckValue(userName string, valueDesired string, valueIndex string) (retBool bool) {
	retBool = false
	us, _ := wapi.ListLocalUsers()
	s := SearchUser(us, userName)
	if valueIndex == "username" {
		if strings.Contains(s.Username, valueDesired) {
			retBool = true
		}
	} else if valueIndex == "fullname" {
		if strings.Contains(s.FullName, valueDesired) {
			retBool = true
		}
	} else if valueIndex == "enabled" {
		if strconv.FormatBool(s.IsEnabled) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "locked" {
		if strconv.FormatBool(s.IsLocked) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "admin" {
		if strconv.FormatBool(s.IsAdmin) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "password_no_expires" {
		if strconv.FormatBool(s.PasswordNeverExpires) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "password_no_change" {
		if strconv.FormatBool(s.NoChangePassword) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "password_age" {
		if strconv.FormatFloat(s.PasswordAge.Hours(), 'F', -1, 64) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "last_login" {
		if s.LastLogon.Format(time.RFC850) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "bad_password" {
		if strconv.FormatUint(uint64(s.BadPasswordCount), 10) == valueDesired {
			retBool = true
		}
	} else if valueIndex == "num_logins" {
		if strconv.FormatUint(uint64(s.NumberOfLogons), 10) == valueDesired {
			retBool = true
		}
	}
	return // return the retBool var
}

//UserCheckExist (This function will tell you first if a record exists in the users list and then allows you to check for a specific value search)
func UserCheckExist(userName string) (retBool bool) {
	retBool = false
	us, _ := wapi.ListLocalUsers()
	retBool = ExistUser(us, userName)
	return // return the retBool var
}

func UserParse(args []string, result interface{}) (retBool bool) {

}