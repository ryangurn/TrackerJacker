package windows

import (
	"TrackerJacker/core"
	"encoding/json"
	"fmt"
	"strings"
)

type ShareStruct []struct {
	Caption        string      `json:"Caption"`
	Description    string      `json:"Description"`
	InstallDate    interface{} `json:"InstallDate"`
	Status         string      `json:"Status"`
	AccessMask     interface{} `json:"AccessMask"`
	AllowMaximum   bool        `json:"AllowMaximum"`
	MaximumAllowed interface{} `json:"MaximumAllowed"`
	Name           string      `json:"Name"`
	Path           string      `json:"Path"`
	Type           int64       `json:"Type"`
}

func ShareExist(share string) (retBool bool) {
	retBool = false

	out := core.Command("Get-WmiObject win32_share | select Name | convertto-json")
	//out := core.Command("Get-WmiObject win32_share | select Caption, Description, InstallDate, Status, AccessMask, AllowMaximum, MaximumAllowed, Name, Path, Type | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if strings.TrimSpace(strings.ToLower(s.Name)) == strings.TrimSpace(strings.ToLower(share)) {
			retBool = true
			return
		}
	}

	return
}

func ShareMeta(share string, key string, value interface{}) (retBool bool) {
	retBool = false

	out := core.Command("Get-WmiObject win32_share | select Caption, Description, InstallDate, Status, AccessMask, AllowMaximum, MaximumAllowed, Name, Path, Type | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if strings.TrimSpace(strings.ToLower(s.Name)) == strings.TrimSpace(strings.ToLower(share)) {
			if key == "Status" {

			} else if key == "Caption" {

			} else if key == "Description" {

			} else if key == "Path" {

			} else if key == "InstallDate" {

			} else if key == "AccessMask" {

			} else if key == "AllowMaximum" {

			} else if key == "MaximumAllowed" {

			} else if key == "Type" {

			}
		}
	}

	return
}

func ShareParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 1 {
		return
	}
	if args[0] == "exist" {
		if ShareExist(args[1]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}