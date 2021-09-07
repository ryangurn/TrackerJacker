package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
)

func ProcessExist(executable string) (retBool bool, retData string) {
	retBool = false

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
			retBool = true
			if out, err := json.Marshal(p); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(processes)
	return retBool, string(out)
}

func ProcessPID(executable string, pid string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
			val, err := strconv.ParseInt(pid, 10, 32)
			if err != nil {
				return
			}

			if p.Pid == int(val) {
				retBool = true
				if out, err := json.Marshal(p); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(processes)
	return retBool, string(out)
}

func ProcessPPID(executable string, ppid string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
			val, err := strconv.ParseInt(ppid, 10, 32)
			if err != nil {
				return
			}

			if p.Ppid == int(val) {
				retBool = true
				if out, err := json.Marshal(p); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(processes)
	return retBool, string(out)
}

func ProcessUsername(executable string, username string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	processes, err := wapi.ProcessList()
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
			if p.Username == username {
				retBool = true
				if out, err := json.Marshal(p); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(processes)
	return retBool, string(out)
}