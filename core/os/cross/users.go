package cross

import (
	"encoding/json"
	"os/user"
)

// TODO: Add support for meta and logged in checks

func UserExist(usr string) (retBool bool, retData string) {
	retBool = false
	u, err := user.Lookup(usr)

	if err != nil {
		return
	}

	if u != nil {
		retBool = true
		if out, err := json.Marshal(u); err == nil {
			retData = string(out)
		}
	}

	return
}