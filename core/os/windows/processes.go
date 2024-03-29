package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
	"strconv"
)

func ProcessExist(executable string) (retBool bool, retData string) {
	retBool = false

	processes, err := wapi.ProcessList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
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
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseInt(pid, 10, 32)
	if err != nil {
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
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
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseInt(ppid, 10, 32)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, p := range processes {
		if p.Executable == executable {
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
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
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