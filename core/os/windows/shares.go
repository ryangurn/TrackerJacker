package windows

import (
	"TrackerJacker/core"
	"encoding/json"
	"strconv"
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

func ShareExist(share string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	out := core.Command("Get-WmiObject win32_share | select Name | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			retBool = true
			if out, err := json.Marshal(s); err == nil {
				return retBool, string(out)
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func ShareStatus(share string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	out := core.Command("Get-WmiObject win32_share | select Name, Status | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.Status == status {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func ShareCaption(share string, caption string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	out := core.Command("Get-WmiObject win32_share | select Name, Caption | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.Caption == caption {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func ShareDescription(share string, description string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	out := core.Command("Get-WmiObject win32_share | select Name, Description | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.Description == description {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func SharePath(share string, path string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	out := core.Command("Get-WmiObject win32_share | select Name, Path | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.Path == path {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func ShareAllowMaximum(share string, maximum string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	val, err := strconv.ParseBool(maximum)
	if err != nil {
		return
	}

	out := core.Command("Get-WmiObject win32_share | select Name, AllowMaximum | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.AllowMaximum == val {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}

func ShareType(share string, typ string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	val, err := strconv.ParseInt(typ, 10, 64)
	if err != nil {
		return
	}

	out := core.Command("Get-WmiObject win32_share | select Name, AllowMaximum | convertto-json")
	var structures ShareStruct
	json.Unmarshal([]byte(out), &structures)
	for _, s := range structures {
		if s.Name == share {
			if s.Type == val {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	o, _ := json.Marshal(structures)
	return retBool, string(o)
}