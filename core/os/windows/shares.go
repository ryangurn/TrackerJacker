package windows

import (
	"TrackerJacker/core"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ShareStruct []struct {
	Caption        string      `json:"Caption"`
	Description    string      `json:"Description"`
	Status         string      `json:"Status"`
	AllowMaximum   bool        `json:"AllowMaximum"`
	Name           string      `json:"Name"`
	Path           string      `json:"Path"`
	Type           int64       `json:"Type"`
}

func ShareExist(share string) (retBool bool) {
	retBool = false

	out := core.Command("Get-WmiObject win32_share | select Name | convertto-json")
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
				if s.Status == value.(string) {
					retBool = true
					return
				}
			} else if key == "Caption" {
				if s.Caption == value.(string) {
					retBool = true
					return
				}
			} else if key == "Description" {
				if s.Description == value.(string) {
					retBool = true
					return
				}
			} else if key == "Path" {
				if s.Path == value.(string) {
					retBool = true
					return
				}
			} else if key == "AllowMaximum" {
				val, err := strconv.ParseBool(value.(string))
				if err != nil {
					return
				}

				if s.AllowMaximum == val {
					retBool = true
					return
				}
			} else if key == "Type" {
				us, err := strconv.ParseInt(value.(string), 10, 64)
				if err != nil {
					return
				}

				if s.Type == us {
					retBool = true
					return
				}
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