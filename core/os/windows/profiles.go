package windows

import (
	"fmt"
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


func ProfileParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 3 {
		return
	}

	if args[0] == "user_directory" {
		if ProfileUserDirectory(args[1]) == result {
			retBool = true
			return
		}
	} else if args[0] == "global_directory" {
		if ProfileDirectory(args[1]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}
