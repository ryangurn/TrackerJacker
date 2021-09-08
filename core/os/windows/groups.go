package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
)

func GroupExist(group string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	groups, err := wapi.ListLocalGroups()
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

	for _, g := range groups {
		if group == g.Name {
			retBool = true
			if out, err := json.Marshal(g); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(groups)
	return retBool, string(out)
}

func GroupComment(group string, comment string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	groups, err := wapi.ListLocalGroups()
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

	for _, g := range groups {
		if group == g.Name {
			if g.Comment == comment {
				retBool = true
				if out, err := json.Marshal(g); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(groups)
	return retBool, string(out)
}