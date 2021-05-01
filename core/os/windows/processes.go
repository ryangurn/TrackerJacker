package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
)

func ProcessExist(processFullPath string) (retBool bool) {
	retBool = false

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Fullpath == processFullPath {
			retBool = true
			return
		}
	}

	return
}

func ProcessMeta(processFullPath string, key string, value interface{}) (retBool bool) {
	retBool = false

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Fullpath == processFullPath {
			if key == "Pid" {
				if p.Pid == value.(int) {
					retBool = true
					return
				}
			} else if key == "Ppid" {
				if p.Ppid == value.(int) {
					retBool = true
					return
				}
			} else if key == "Username" {
				if p.Username == value.(string) {
					retBool = true
					return
				}
			} else if key == "Executable" {
				if p.Executable == value.(string) {
					retBool = true
					return
				}
			}
		}
	}

	return
}