package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	"github.com/ceshihao/windowsupdate"
	"github.com/go-ole/go-ole"
	wapi "github.com/iamacarpet/go-win64api"
	"github.com/scjalliance/comshim"
	"os"
)

func getUpdateHistory() (data []*windowsupdate.IUpdateHistoryEntry) {
	comshim.Add(1)
	defer comshim.Done()

	// ole.CoInitialize(0)
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	defer ole.CoUninitialize()

	var err error

	session, err := windowsupdate.NewUpdateSession()
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
	}

	// Query Update History
	searcher, err := session.CreateUpdateSearcher()
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
	}

	result, err := searcher.QueryHistoryAll()
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
	}

	return result
}

func UpdateCompleted(update string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history := getUpdateHistory()
	for _, h := range history {
		if h.Title == update {
			retBool = true
			if out, err := json.Marshal(h); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(history)
	return retBool, string(out)
}

func UpdatePending(update string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history, err := wapi.UpdatesPending()
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

	for _, h := range history.UpdateHistory {
		if h.UpdateName == update {
			retBool = true
			if out, err := json.Marshal(h); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(history)
	return retBool, string(out)
}

func UpdateHistory(update string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	history, err := wapi.UpdatesPending()
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

	for _, h := range history.UpdateHistory {
		if h.UpdateName == update {
			if h.Status == status {
				retBool = true
				if out, err := json.Marshal(h); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(history)
	return retBool, string(out)
}