package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
)

func SoftwareExist(software string) (retBool bool){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			retBool = true
		}
	}

	return
}

func SoftwareMeta(software string, key string, value interface{}) (retBool bool) {
	retBool = false

	return
}

func SoftwareParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 1 {
		return
	}

	if args[0] == "exist" {
		if len(args) != 2 {
			return
		}

		if SoftwareExist(args[1]) == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}
