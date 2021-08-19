package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func FirewallEnabled(profile string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	val, err := strconv.ParseInt(profile, 10, 32)
	if err != nil {
		return
	}

	res, err := wapi.FirewallIsEnabled(int32(val))
	if err != nil {
		return
	}

	retBool = res
	if out, err := json.Marshal(res); err == nil {
		return retBool, string(out)
	}

	return
}

func FirewallProfile(firewall string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	profiles, err := wapi.FirewallCurrentProfiles()
	if err != nil {
		return
	}

	switch firewall {
	case "Public":
		retBool = profiles.Public
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	case "Domain":
		retBool = profiles.Domain
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	case "Private":
		retBool = profiles.Private
		if out, err := json.Marshal(profiles); err == nil {
			return retBool, string(out)
		}
		break
	}

	return
}