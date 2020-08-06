package windows

import (
	"fmt"
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

func ProcessParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) < 2 {
		return
	}

	if args[0] == "exist" {
		if len(args) != 2 {
			return
		}

		if ProcessExist(args[1]) == result {
			retBool = true
			return
		}
	} else if args[0] == "meta" {
		if len(args) != 4 {
			return
		}

		if ProcessMeta(args[1], args[2], args[3]) == result {
			retBool = true
			return
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}
