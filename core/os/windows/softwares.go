package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
)

func SoftwareExist(software string) (retBool bool){
	retBool = false

	return
}

func SoftwareMeta(software string, key string, value interface{}) (retBool bool) {
	retBool = false

	return
}

func SoftwareParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	sw, _ := wapi.InstalledSoftwareList()

	for s, k := range sw{
		fmt.Println(s, k.)
	}

	return
}
