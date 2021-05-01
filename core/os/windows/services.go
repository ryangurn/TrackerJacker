package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"strings"
)

func ServiceExist(svc string) (retBool bool) {
	retBool = false

	services, err := wapi.GetServices()
	if err != nil {
		return
	}

	for _, s := range services {
		if strings.TrimSpace(strings.ToLower(s.SCName)) == strings.TrimSpace(strings.ToLower(svc)) {
			retBool = true
			return
		}
	}

	return
}

func ServiceMeta(svc string, key string, value interface{}) (retBool bool) {
	retBool = false

	services, err := wapi.GetServices()
	if err != nil {
		return
	}

	for _, s := range services {
		if strings.TrimSpace(strings.ToLower(s.SCName)) == strings.TrimSpace(strings.ToLower(svc)) {
			if key == "DisplayName" {
				if s.DisplayName == value.(string) {
					retBool = true
					return
				}
			} else if key == "StatusText" {
				if s.StatusText == value.(string) {
					retBool = true
					return
				}
			} else if key == "Status" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if s.Status == uint32(us) {
					retBool = true
					return
				}
			} else if key == "AcceptStop" {
				b, err := strconv.ParseBool(value.(string))
				if err != nil {
					return
				}

				if s.AcceptStop == b {
					retBool = true
					return
				}
			} else if key == "IsRunning" {
				b, err := strconv.ParseBool(value.(string))
				if err != nil {
					return
				}

				if s.IsRunning == b {
					retBool = true
					return
				}
			} else if key == "RunningPid" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if s.RunningPid == uint32(us) {
					retBool = true
					return
				}
			} else if key == "ServiceType" {
				us, err := strconv.ParseUint(value.(string), 10, 32)
				if err != nil {
					return
				}

				if s.ServiceType == uint32(us) {
					retBool = true
					return
				}
			}
		}
	}

	return
}