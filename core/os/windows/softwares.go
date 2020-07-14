package windows

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"time"
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

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil{
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software{
			if key == "DisplayVersion" {
				if v.DisplayVersion == value.(string) {
					retBool = true
					return
				}
			} else if key == "Arch" {
				if v.Arch == value.(string) {
					retBool = true
					return
				}
			} else if key == "Publisher" {
				if v.Publisher == value.(string){
					retBool = true
					return
				}
			} else if key == "InstallDate" {
				if v.InstallDate == value.(time.Time) {
					retBool = true
					return
				}
			} else if key == "EstimatedSize" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if v.EstimatedSize == us {
					retBool = true
					return
				}
			} else if key == "Contact" {
				if v.Contact == value.(string) {
					retBool = true
					return
				}
			} else if key == "HelpLink" {
				if v.HelpLink == value.(string) {
					retBool = true
					return
				}
			} else if key == "InstallSource" {
				if v.InstallSource == value.(string) {
					retBool = true
					return
				}
			} else if key == "InstallLocation" {
				if v.InstallLocation == value.(string) {
					retBool = true
					return
				}
			} else if key == "VersionMajor" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if v.VersionMajor == us {
					retBool = true
					return
				}
			} else if key == "VersionMinor" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if v.VersionMinor == us {
					retBool = true
					return
				}
			}
		}
	}

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
	} else if args[0] == "meta" {
		if len(args) != 4 {
			return
		}
		if SoftwareMeta(args[1], args[2], args[3]) == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}
