package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
)

func ProfileUserDirectory(dir string) (retBool bool) {
	retBool = false

	profile, err := wapi.GetDefaultUserProfileDirectory()
	if err != nil{
		return
	}

	if profile == dir {
		retBool = true
		return
	}

	return
}

func ProfileDirectory(dir string) (retBool bool) {
	retBool = false

	profile, err := wapi.GetProfilesDirectory()
	if err != nil {
		return
	}

	if profile == dir {
		retBool = true
		return
	}

	return
}
