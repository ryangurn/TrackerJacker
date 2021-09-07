package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
)

func ProfileUserDirectory(dir string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profile, err := wapi.GetDefaultUserProfileDirectory()
	if err != nil{
		return
	}

	if profile == dir {
		retBool = true
		if out, err := json.Marshal(profile); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(profile)
	return retBool, string(out)
}

func ProfileDirectory(dir string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profile, err := wapi.GetProfilesDirectory()
	if err != nil {
		return
	}

	if profile == dir {
		retBool = true
		if out, err := json.Marshal(profile); err == nil {
			return retBool, string(out)
		}
	}

	out, _ := json.Marshal(profile)
	return retBool, string(out)
}
